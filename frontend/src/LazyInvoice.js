import { hot } from 'react-hot-loader/root';
import React from 'react';
import { Query } from 'react-apollo';
import gql from 'graphql-tag';
import { format } from 'date-fns';

const GET_INVOICES = gql`
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
`;

export const Component = () =>
  <div>
    <div data-e2e="example">
      Hello Feibian!
    </div>
    <Query query={GET_INVOICES}>
      {({ data, loading, error }) => {
        if (loading) return <div>Loading ...</div>;
        if (error) return <div>Error!</div>;

        return (
          <div>
            {data.invoices.map((invoice, i) =>
              <div key={i}>
                {`Invoice date: ${format(invoice.date, 'd MMMM, YYYY' )}`}
                {invoice.items.map((item, i) =>
                  <div key={i} style={{ display: 'flex'}}>
                    <div style={{ width: 150 }}>{item.description}</div>
                    <div style={{ width: 50 }}>{item.quantity}</div>
                    <div style={{ width: 50 }}>{item.unitPrice}</div>
                  </div>
                )}
              </div>)}
          </div>
        );
      }}
    </Query>
  </div>;

export default hot(Component);
