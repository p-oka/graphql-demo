import React from "react";
import Grid from '@material-ui/core/Grid';
import ApolloClient from "apollo-boost";
import { ApolloProvider } from "react-apollo";

import AddTask from "./AddTask";
import TaskList from "./TaskList";

export default function App() {
  return (
    <ApolloProvider client={new ApolloClient()}>
      <Grid container >
        <Grid item xs={2} />
        <Grid item xs={8}>
          <AddTask />
          <TaskList />
        </Grid>
        <Grid item xs={2} />
      </Grid>
    </ApolloProvider>
  );
}
