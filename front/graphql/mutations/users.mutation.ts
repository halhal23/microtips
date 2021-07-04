import { gql } from '@apollo/client';
// import { User } from '../../types'

export const SIGN_IN_MUTATION = gql`
  mutation SignIn($name: String!, $password: String!) {
    SignIn(input: {name: $name, password: $password})
  }
`;