import { ApolloClient, InMemoryCache, ApolloProvider, HttpLink } from '@apollo/client';
import { AppProps } from 'next/dist/next-server/lib/router/router';

const link = new HttpLink({
  uri: 'http://localhost:8080/query',
})

const cache = new InMemoryCache();
const client = new ApolloClient({
  link,
  cache
})

function MyApp({Component, PageProps}: AppProps) {
  return (
    <ApolloProvider client={client}>
      <Component {...PageProps} />
    </ApolloProvider>
  )
}

export default MyApp;
