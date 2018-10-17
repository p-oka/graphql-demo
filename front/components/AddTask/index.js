import React from "react";
import { Mutation } from "react-apollo";
import Grid from '@material-ui/core/Grid';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';

import enhance from "./enhance";
import { GET_TASKS } from "/graphql/queries";
import { CREATE_TASK } from "/graphql/mutations";

class AddTask extends React.Component {
  constructor(props) {
    super(props);
    this.nameFieldRef = React.createRef();
  }

  render() {
    const { classes } = this.props;

    return (
      <Mutation
        mutation={CREATE_TASK}
        refetchQueries={() => [{ query: GET_TASKS }]}
      >
        {createTask => (
          <Grid container spacing={8}>
            <Grid item xs={10}>
              <TextField
                id="name"
                label="やること"
                margin="normal"
                fullWidth
                inputRef={this.nameFieldRef}
              />
            </Grid>
            <Grid className={classes.centerd} item xs={2}>
              <Button
                variant="contained"
                color="primary"
                fullWidth
                onClick={() => {
                  const node = this.nameFieldRef.current;
                  createTask({ variables: { name: node.value }});
                  node.value = "";
                }}
              >
                作成
              </Button>
            </Grid>
          </Grid>
        )}
      </Mutation>
    );
  }
}

export default enhance(AddTask);