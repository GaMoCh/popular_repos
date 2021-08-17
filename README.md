<div align="center">
  <h2>Popular Repos</h2>
</div>

<div align="center">
  <a href="#about">About</a>
   &nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#techs">Techs</a>
  &nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
    <a href="#techs">Architecture</a>
  &nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
    <a href="#features">Features</a>
  &nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#run">Run</a>
&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#license">License</a>
</div>

## About

In this work, we will study the main characteristics of popular open-source systems. Thus, we will analyze how they are developed, how often they receive external input, how often they release releases, among other characteristics. To do so, we will collect the data indicated below for the 1,000 repositories with the highest number of stars on GitHub and discuss the values obtained.

## Techs

- Golang `1.16`

## Architecture

We based on project layout [golang-standards/project-layout](https://github.com/golang-standards/project-layout)

## Features

- [Lab01S01](https://github.com/gamoch/popular_repos/releases/tag/v0.1.1): GraphQL query for 100 repositories + automatic request

- Lab01S02: Pagination (query for 1000 repositories) + data in `.csv` file

- Lab01S03: Data analysis and visualization + final report

## Run

1. Go to [releases](https://github.com/gamoch/popular_repos/releases)

2. Do download based on your Operational System and Architecture

3. Extract the archive

4. Run the executable inside `./bin` and provide your GitHub token with flag `-token` or use environment variable `GITHUB_TOKEN`

5. Enjoy it!

## License

This project is under the MIT license. See the file [LICENSE](LICENSE) for more details.
