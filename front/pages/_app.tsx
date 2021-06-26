import { ApolloClient, InMemoryCache, ApolloProvider } from '@apollo/client';
import { AppProps } from 'next/dist/next-server/lib/router/router';

const cache = new InMemoryCache();
const client = new ApolloClient({
  uri: 'http://localhost:8080/query',
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
