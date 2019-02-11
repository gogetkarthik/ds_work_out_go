package tips

//.gitignore https://www.atlassian.com/git/tutorials/saving-changes/gitignore

// You can also define personal ignore patterns for a particular repository in a special file at .git/info/exclude

//In addition, you can define global Git ignore patterns for all repositories on your local system by setting the Git core.excludesFile property.

//You can use the git check-ignore command with the -v (or --verbose) option to determine which pattern is causing a particular file to be ignored:

//git customizing https://git-scm.com/book/en/v1/Customizing-Git-Git-Configuration

//Git looks for these values is in an /etc/gitconfig file, which contains values for every user on the system and all of their repositories. If you pass the option --system to git config, it reads and writes from this file specifically.

//The next place Git looks is the ~/.gitconfig file, which is specific to each user. You can make Git read and write to this file by passing the --global option.

//Finally, Git looks for configuration values in the config file in the Git directory (.git/config) of whatever repository you’re currently using.

//commit.template  -- suppose you create a template file at $HOME/.gitmessage.txt that looks like this:
