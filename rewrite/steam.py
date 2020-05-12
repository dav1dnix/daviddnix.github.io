import requests
import os
from dotenv import load_dotenv
load_dotenv()

class User:
    def __init__(self, key, id):
        self.key = key
        self.id = id
        self.getreq = requests.get(f"https://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/?key={self.key}&steamids={self.id}").json()


    def api_key(self):
        if not self.key:
            return "You didn't provide an API key"
        else:
            return

    def json_response(self):
        return self.getreq["response"]["players"][0]

    def profile(self):
        r = self.json_response()
        Username = f"Username: {r['personaname']}"
        ID = f"ID: {r['steamid']}"
        Profile = f"Profile: {r['profileurl']}"
        PFP = f"PFP: {r['avatar']}"
        Country = f"Country: {r['loccountrycode']}"
        print(Username, ID, Profile, PFP, Country, sep="\n")

obj = User(os.getenv("API_KEY"), "76561198353821873")
obj.profile()