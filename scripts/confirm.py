# This script prompts the user to enter 
# nN (no) or yY (yes). If yes is provided
# it runs the command in sys.argv[1].
#
# It is meant to be used for confirming
# operations in makefiles.
import sys
from subprocess import call

v = raw_input("Are you sure? (N,y): ")
if v.lower() == "y":
    commands = sys.argv[1].split("&&")
    for cmd in commands:
        call(cmd.strip().split(" "))