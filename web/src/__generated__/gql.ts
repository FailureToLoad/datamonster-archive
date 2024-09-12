/* eslint-disable */
import * as types from './graphql';
import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';

/**
 * Map of all GraphQL operations in the project.
 *
 * This map has several performance disadvantages:
 * 1. It is not tree-shakeable, so it will include all operations in the project.
 * 2. It is not minifiable, so the string of a GraphQL query will be multiple times inside the bundle.
 * 3. It does not support dead code elimination, so it will add unused operations.
 *
 * Therefore it is highly recommended to use the babel or swc plugin for production.
 */
const documents = {
    "\n  mutation CreateSurvivor($input: CreateSurvivorInput!) {\n    createSurvivor(input: $input) {\n      id\n      name\n    }\n  }\n": types.CreateSurvivorDocument,
    "\n  mutation UpdateSurvivor($id: ID!, $input: UpdateSurvivorInput!) {\n    updateSurvivor(id: $id, input: $input) {\n      id\n      name\n    }\n  }\n": types.UpdateSurvivorDocument,
    "\n  mutation DeleteSurvivor($id: ID!) {\n    deleteSurvivor(id: $id)\n  }\n": types.DeleteSurvivorDocument,
    "\n  query GetSurvivors($settlementId: ID!) {\n    survivors(filter: {settlementID: $settlementId}) {\n      id\n      accuracy\n      born\n      courage\n      evasion\n      gender\n      huntxp\n      insanity\n      luck\n      lumi\n      movement\n      name\n      speed\n      strength\n      survival\n      systemicpressure\n      torment\n      understanding\n    }\n  }\n": types.GetSurvivorsDocument,
    "\n  mutation CreateSettlement($input: CreateSettlementInput!) {\n    createSettlement(input: $input) {\n      id\n      name\n      owner\n    }\n  }\n": types.CreateSettlementDocument,
    "\n  query GetSettlements {\n    settlements {\n      id\n      name\n    }\n  }\n": types.GetSettlementsDocument,
};

/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 *
 *
 * @example
 * ```ts
 * const query = gql(`query GetUser($id: ID!) { user(id: $id) { name } }`);
 * ```
 *
 * The query argument is unknown!
 * Please regenerate the types.
 */
export function gql(source: string): unknown;

/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  mutation CreateSurvivor($input: CreateSurvivorInput!) {\n    createSurvivor(input: $input) {\n      id\n      name\n    }\n  }\n"): (typeof documents)["\n  mutation CreateSurvivor($input: CreateSurvivorInput!) {\n    createSurvivor(input: $input) {\n      id\n      name\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  mutation UpdateSurvivor($id: ID!, $input: UpdateSurvivorInput!) {\n    updateSurvivor(id: $id, input: $input) {\n      id\n      name\n    }\n  }\n"): (typeof documents)["\n  mutation UpdateSurvivor($id: ID!, $input: UpdateSurvivorInput!) {\n    updateSurvivor(id: $id, input: $input) {\n      id\n      name\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  mutation DeleteSurvivor($id: ID!) {\n    deleteSurvivor(id: $id)\n  }\n"): (typeof documents)["\n  mutation DeleteSurvivor($id: ID!) {\n    deleteSurvivor(id: $id)\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query GetSurvivors($settlementId: ID!) {\n    survivors(filter: {settlementID: $settlementId}) {\n      id\n      accuracy\n      born\n      courage\n      evasion\n      gender\n      huntxp\n      insanity\n      luck\n      lumi\n      movement\n      name\n      speed\n      strength\n      survival\n      systemicpressure\n      torment\n      understanding\n    }\n  }\n"): (typeof documents)["\n  query GetSurvivors($settlementId: ID!) {\n    survivors(filter: {settlementID: $settlementId}) {\n      id\n      accuracy\n      born\n      courage\n      evasion\n      gender\n      huntxp\n      insanity\n      luck\n      lumi\n      movement\n      name\n      speed\n      strength\n      survival\n      systemicpressure\n      torment\n      understanding\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  mutation CreateSettlement($input: CreateSettlementInput!) {\n    createSettlement(input: $input) {\n      id\n      name\n      owner\n    }\n  }\n"): (typeof documents)["\n  mutation CreateSettlement($input: CreateSettlementInput!) {\n    createSettlement(input: $input) {\n      id\n      name\n      owner\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query GetSettlements {\n    settlements {\n      id\n      name\n    }\n  }\n"): (typeof documents)["\n  query GetSettlements {\n    settlements {\n      id\n      name\n    }\n  }\n"];

export function gql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> = TDocumentNode extends DocumentNode<  infer TType,  any>  ? TType  : never;