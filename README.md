# Seedr

![Seedr Logo](https://avatars3.githubusercontent.com/u/54958814?s=200&v=4)

Seedr helps you to get started quickly. Seedr can setup a project skeleton for 
you in seconds. 

## How does it work

You give `seedr` a seed, that can either be a git repository, a directory or a 
predefined seed name. With information, seedr ask you some more questions and 
build your skeleton project.

If the seed contains a `Seedfile.yml` some more questions will be asked. 
These will customize the skeleton to save you some more time.


## Getting started

Download the latest version:

```bash
curl -sL http://bit.ly/gh-get | PROJECT=sourceseed/seedr bash
```

Generate your first project:

```
seedr generate --seed golang --target test123
```


## Seeds

 - directory on your local file system
 - git repository (only public at this point)
 - predefined seed, will fetch the seed from `github.com/sourceseed/seed-<NAME>.git`


### Seedfile.yml

```yaml
name: nameOfSeed
parameters:
  - variable: APPNAME
    description: "Application name"
```

If a seed contains a `Seedfile.yml`, values for all parameters will be requested
during setup. Seedr will search for `__APPNAME__` and replace it with the 
provide value in both file name and file content.


### .seedkeep

If you want to keep empty directories in your seed you can place `.seedkeep`
files in these directories. These will be deleted upon generation allowing for
later usage.


## Roadmap

 - Cleanup output
 - Allow non-interactive generating
 - Add parameter validation options
