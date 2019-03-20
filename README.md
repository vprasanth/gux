# gux

gux is yet another tmux wrapper with a user centric focus on simplifying multiple tty workflows, specifically around shareability, standardization, and easy of use. 

## Status
_In development_

## Goals

- define a tmux window/pane template specification (.gux file)
- shareable templates (fetchable)
- be SUPER easy & straightforward
- provide extensible API 

### .gux spec (_in development_)
```yaml
version: "0.1.0"
session: 
  - name: test 
    window:
      - layout: vertical-split
        name: editor 
        workingDir: ~/code
        panes:
          - command: vim .
          - command: bash
          - command: figlet hello
    workingDir: ~/code
```
