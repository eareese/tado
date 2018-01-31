# About this project

Create, list, and do (complete and delete) todo list tasks. Data is stored using BoltDB.

Use `todo check` to hit the Canvas API. This will query for all ungraded assignments and prompt the user to add tasks of the format "Grade [assignment name]" to the todo list. It requires an API access token to be present in an environment variable called `TODO_TOKEN`.
