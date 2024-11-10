import {Survivor} from '@/lib/types/survivor';

const updateUrl = (settlementId: string, survivorId: number) => {
  return `${
    import.meta.env.VITE_API_HOST
  }/settlements/${settlementId}/survivors/${survivorId}`;
};

const insertUrl = (settlementId: string) => {
  return `${
    import.meta.env.VITE_API_HOST
  }/settlements/${settlementId}/survivors`;
};

export async function CreateSurvivor(
  survivor: Survivor,
  settlementId: string,
  accessToken: string
) {
  const response = await fetch(insertUrl(settlementId), {
    method: 'post',
    headers: new Headers({
      Authorization: `Bearer ${accessToken}`,
      'Content-Type': 'application/json',
    }),
    body: JSON.stringify(survivor),
  });

  if (!response.ok) {
    throw new Error('survivor creation failed');
  }
}

export async function UpdateSurvivor(
  survivor: Survivor,
  settlementId: string,
  accessToken: string
) {
  const response = await fetch(updateUrl(settlementId, survivor.id), {
    method: 'post',
    headers: new Headers({
      Authorization: `Bearer ${accessToken}`,
      'Content-Type': 'application/json',
    }),
    body: JSON.stringify(survivor),
  });

  if (!response.ok) {
    throw new Error('survivor creation failed');
  }
}

export async function FetchSurvivors(
  settlementId: string,
  accessToken: string
) {
  const path = `${
    import.meta.env.VITE_API_HOST
  }/settlements/${settlementId}/survivors`;
  console.log(`request path ${path}`);
  const response = await fetch(path, {
    method: 'get',
    headers: new Headers({
      Authorization: `Bearer ${accessToken}`,
      'Content-Type': 'application/json',
    }),
  });

  if (!response.ok) {
    throw new Error('unable to fetch survivors');
  }

  const responseJson: Array<Survivor> = await response.json();
  return responseJson;
}
