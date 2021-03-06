import { gql } from "apollo-boost";
import { graphql } from 'react-apollo'
import { compose, withProps, branch, renderComponent } from "recompose";

import LoadingScene from "/components/LoadingScene";
import { GET_TASKS } from "/graphql/queries";


const withQuery = graphql(GET_TASKS);

const withLoadingScene = branch(
  props => props.data.networkStatus <= 1,
  renderComponent(LoadingScene)
);

const passProps = withProps(
  props => {
    const { tasks } = props.data;
    return { tasks };
  }
);

export default compose(
  withQuery,
  withLoadingScene,
  passProps
);