import React from 'react';
import ReactDOM from 'react-dom';
import gql from 'graphql-tag';
import { ApolloClient } from 'apollo-client';
import { ApolloProvider } from 'react-apollo';
import { HttpLink } from 'apollo-link-http';
import { InMemoryCache } from 'apollo-cache-inmemory';

import LazyInvoice from './LazyInvoice';

const cache = new InMemoryCache();

const link = new HttpLink({
  uri: 'http://localhost:8081/query',
});

const client = new ApolloClient({
  cache,
  link
});

client
  .query({
    query: gql`
      query GetInvoices {
        invoices {
          date
          items {
            description
            quantity
            unitPrice
            }
          }
        }
      `
  })
  .then(result => console.log(result)); // eslint-disable-line no-console

ReactDOM.render(
  <ApolloProvider client={client}>
    <LazyInvoice />
  </ApolloProvider>,
  document.getElementById('root')
);
