# datatug-github-action

A GitHub action for [DataTug](https://datatug.app/) projects.


## Usage

Add following step to the end of your workflow:

    - uses: datatug/datatug-github-action
      if: always()
        with:
          datatug_projects: <DATATUG_PROJECT_FOLDER> 

The `datatug_projects` is optional and set to `"datatug"` by default. 


## [License](LICENSE)

Free & open source under MIT license

&copy; 2021 [Sneat.team](https://sneat.team)
