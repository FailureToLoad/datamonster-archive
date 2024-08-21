export type Settlement = SettlementId & {
  limit: number;
  departing: number;
  cc: number;
  year: number;
};

export type SettlementId = {
  id: string;
  name: string;
};
