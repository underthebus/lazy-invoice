import { Selector } from 'testcafe';

fixture `Getting Started`
  .page `http://localhost:8080`;

test('Test home page renders in dev environment', async t => {
  const exampleEl = Selector('[data-e2e="example"]', { timeout: 5000 });
  await t.expect(exampleEl.innerText).eql('Hello Feibian!');
});
