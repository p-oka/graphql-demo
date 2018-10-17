import { gql } from "apollo-boost";

export const CREATE_TASK = gql`
  mutation CreateTask($name: String!) {
    createTask(name: $name) {
      id
      name
    }
  }
`;