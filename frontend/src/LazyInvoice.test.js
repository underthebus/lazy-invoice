import React from 'react';
import { shallow } from 'enzyme';

import { Component } from './LazyInvoice';

it('renders as expected', () => {
  const wrapper = shallow(<Component />);
  expect(wrapper).toMatchSnapshot();
});
