# Tally backend take-home

Hi! Thanks for your interest in working at Tally. The first step in our process is this short
back-end coding task. We expect this task to take a few hours.

We’re looking for a solution that is both easy-to-read and correct. We will focus on clean
code and correctness of the business logic.

We don’t expect you to handle all the errors or edge cases. Focus on making sure that
the happy path works. We are most interested in the schema, architecture, and external
integrations.


# The problem

You would like to get notified by email whenever there is a new ENS DAO proposal to vote on.

ENS is a name system that runs on Ethereum smart contracts. ENS token holders control ENS
via a Governor smart contract. Tokenholders can propose changes to ENS by submitting
proposals to the Governor. If a vote passes, the contract implements the proposed changes.

But Ethereum cannot send email! You’ll need to run something off chain to send the email.

Tally API

Tally’s indexer already indexes ENS DAO proposals from the blockchain. You can fetch them
with our public API.

Here’s an overview of the Tally API: https://docs.tally.xyz/tally-api/welcome. There’s a link to the
API docs and instructions for getting an API key on that page.

# Requirements

Build a service to send yourself an email when there is a new ENS proposal. The body of the
email should contain the title of the proposal and a link to it on Tally.

To keep things simple, you don’t need to handle multiple users..

# Solution

Please implement the solution in Go. Include a clear README with instructions for testing your
service.

When you’re finished, email us to let us know. Please don’t post it in a public git repository, as
we may reuse this question in the future. You’re welcome to use AI coding tools, as long as you
know what your code does!

Let us know if you have any questions or if something does not seem to make sense
