#!/usr/bin/osascript

on run (args)
  display dialog (first item of args) with title (second item of args) with icon file (third item of args) buttons {"What?!", "Thanks, I appreciate it."} giving up after 5
end run