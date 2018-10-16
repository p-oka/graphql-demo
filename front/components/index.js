import React from "react";
import { ApolloProvider } from "react-apollo";
import ApolloClient from "apollo-boost";

import TaskList from "./TaskList";

export default function App() {
  return (
    <ApolloProvider client={new ApolloClient()}>
      <TaskList />
    </ApolloProvider>
  );
}
