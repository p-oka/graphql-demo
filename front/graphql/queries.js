import { gql } from "apollo-boost";

export const GET_TASKS = gql`
  query GetTasks {
    tasks {
      id
      name
    }
  }
`;