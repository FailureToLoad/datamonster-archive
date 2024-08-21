import { render, screen, waitFor } from "@testing-library/react";
import { describe, it, expect, vi, Mock } from "vitest";
import userEvent from "@testing-library/user-event";
import { CreateSettlementDialog } from "./creationDialog";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

function renderComponent(
  queryClient: QueryClient,
  mockCreateSettlement: Mock,
  mockGetToken: Mock
) {
  return render(
    <QueryClientProvider client={queryClient}>
      <CreateSettlementDialog
        createSettlement={mockCreateSettlement}
        getToken={mockGetToken}
      />
    </QueryClientProvider>
  );
}

describe("Settlements Modal", () => {
  const queryClient = new QueryClient({
    defaultOptions: {
      queries: {
        retry: false,
      },
    },
  });
  const mockCreateSettlement = vi.fn();
  const mockGetToken = vi.fn();
  it("should launch the modal when the plus button is clicked", async () => {
    const { getByLabelText } = renderComponent(
      queryClient,
      mockCreateSettlement,
      mockGetToken
    );
    const plusButton = getByLabelText("Create Settlement");
    expect(plusButton).toBeInTheDocument();
    userEvent.click(plusButton);

    await waitFor(() => {
      expect(screen.queryByTestId("settlement-modal")).toBeInTheDocument();
    });
  });
  it("should not submit short settlement names", async () => {
    const { getByLabelText } = renderComponent(
      queryClient,
      mockCreateSettlement,
      mockGetToken
    );
    const plusButton = getByLabelText("Create Settlement");
    expect(plusButton).toBeInTheDocument();
    const user = userEvent.setup();
    user.click(plusButton);

    await vi.waitFor(() => {
      expect(screen.queryByTestId("settlement-modal")).toBeInTheDocument();
    });

    const submitButton = screen.getByRole("button", { name: /add/i });
    user.click(submitButton);

    const nameInput = screen.getByLabelText("Settlement Name");
    await vi.waitFor(() => {
      expect(nameInput.getAttribute("aria-invalid")).toBe("true");
    });
    expect(mockCreateSettlement).not.toHaveBeenCalled();
  });
  it("should not submit long settlement names", async () => {
    const { getByLabelText } = renderComponent(
      queryClient,
      mockCreateSettlement,
      mockGetToken
    );
    const plusButton = getByLabelText("Create Settlement");
    expect(plusButton).toBeInTheDocument();
    const user = userEvent.setup();
    user.click(plusButton);

    await vi.waitFor(() => {
      expect(screen.queryByTestId("settlement-modal")).toBeInTheDocument();
    });

    const nameInput = screen.getByLabelText("Settlement Name");

    const tooLongText = "aaaaAaaaaAaaaaAaaaaAaaaaAaaaaA";
    user.click(nameInput);
    user.paste(tooLongText);

    const submitButton = screen.getByRole("button", { name: /add/i });
    user.click(submitButton);
    await vi.waitFor(() => {
      expect(nameInput.getAttribute("aria-invalid")).toBe("true");
    });
    expect(mockCreateSettlement).not.toHaveBeenCalled();
  });
  it("should submit valid names", async () => {
    const { getByLabelText } = renderComponent(
      queryClient,
      mockCreateSettlement,
      mockGetToken
    );
    const plusButton = getByLabelText("Create Settlement");
    expect(plusButton).toBeInTheDocument();
    const user = userEvent.setup();
    user.click(plusButton);

    await vi.waitFor(() => {
      expect(screen.queryByTestId("settlement-modal")).toBeInTheDocument();
    });

    // Find the input field for the settlement name
    const nameInput = screen.getByLabelText("Settlement Name");
    user.click(nameInput);
    user.paste("valid");
    // Enter a short settlement name
    // Assert that the input field has aria-invalid set to true
    await vi.waitFor(() => {
      expect(nameInput.getAttribute("aria-invalid")).toBe("false");
    });
    // Click the "Add" button to submit the form
    const submitButton = screen.getByRole("button", { name: /add/i });
    user.click(submitButton);

    await vi.waitFor(() => {
      expect(mockCreateSettlement).toHaveBeenCalled();
    });
  });
});
