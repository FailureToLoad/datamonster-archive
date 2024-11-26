export type Survivor = {
  id: number;
  settlementID: number;
  name: string;
  born: number;
  gender: 'M' | 'F';
  status: 'alive' | 'dead' | 'retired';
  huntXp: number;
  survival: number;
  movement: number;
  accuracy: number;
  strength: number;
  evasion: number;
  luck: number;
  speed: number;
  insanity: number;
  systemicPressure: number;
  torment: number;
  lumi: number;
  courage: number;
  understanding: number;
};

export const DefaultSurvivor: Survivor = {
  id: 0,
  settlementID: 0,
  name: 'Meat',
  born: 0,
  gender: 'M',
  huntXp: 0,
  survival: 1,
  systemicPressure: 0,
  movement: 5,
  accuracy: 0,
  strength: 0,
  evasion: 0,
  luck: 0,
  speed: 0,
  lumi: 0,
  insanity: 0,
  torment: 0,
  courage: 0,
  understanding: 0,
  status: 'alive',
};

export enum SurvivorStatus {
  Alive = 'ALIVE',
  CeasedToExist = 'CEASED_TO_EXIST',
  Dead = 'DEAD',
  Retired = 'RETIRED',
  SkipHunt = 'SKIP_HUNT',
}
