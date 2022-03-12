import webbrowser
import time

lst = ["Sand foundary", "w3 Adda", "w3spoit", "devhints.io", "bootstrap cheatsheet",
       "google lighthouse", "gitignore.io", "meta tags", "frontendchecklist.io",
       "regex101", "SWdaily", "Alison", "Thinkful", "app academy", "Kenzie Academy",
       "HTML Dog", "Dev.to", "Swift.com", "SkillShare", "Goodle Devs", "AppCoda", "TreeHouse",
       "Ray Wenderlich", "FutureLearn", "NextGenT", "CISA", "OpenSecurity", "Heimdal", "Sans Course",
       "DataQuest", "DataSchool", "Heroicons", "Feather Icons", "Css.gg", "Ionicons", "Font Awesome", "Flaticon",
       "weworkremotely", "virtualvocations", "flexjobs", "justremote", "remote", "photopea", "ray.so", "jitter.video",
       "polarite.app", "attentioninsight.com", "Mubert", "HighTouch", "CVshine", "TwittrGems", "hidden tools",
       "Can i use", "BundlePhobia", "CodeElf"]


for x in lst:
    webbrowser.open('https://www.google.com/search?q='+x, new=1)
    time.sleep(0.5)
