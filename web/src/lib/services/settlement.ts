import { translateFormError } from "@/lib/utils";
import { Settlement } from "@/lib/types/settlements";

type SettlementCreationRequest = {
  name: string;
};
export const CreateSettlement = async (
  settlementName: string,
  accessToken: string
) => {
  try {
    const request: SettlementCreationRequest = {
      name: settlementName,
    };
    const response = await fetch(
      `${import.meta.env.VITE_API_HOST}/settlements`,
      {
        method: "post",
        headers: new Headers({
          Authorization: "Bearer " + accessToken,
          "Content-Type": "application/json",
        }),
        body: JSON.stringify(request),
      }
    );
    if (!response.ok) {
      throw new Error("unable to create settlement");
    }
  } catch (error) {
    throw new Error(translateFormError(error));
  }
};

export async function GetSettlements(
  accessToken: string
): Promise<Settlement[]> {
  const response = await fetch(`${import.meta.env.VITE_API_HOST}/settlements`, {
    method: "get",
    headers: new Headers({
      Authorization: `Bearer ${accessToken}`,
      "Content-Type": "application/json",
    }),
  });

  if (!response.ok) {
    throw new Error("unable to fetch settlements for user");
  }

  const responseJson: Array<Settlement> = await response.json();
  return responseJson;
}
