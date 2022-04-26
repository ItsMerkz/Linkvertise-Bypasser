import httpx

class linkbypass():
    def __init__(self):
        self.url = input("[+] Url > ")

    def bypass(self):
        resp = httpx.get(f"https://bypass.bot.nu/bypass2?url={self.url}")
        self.timetaken = resp.json()["time_ms"]
        self.bypassedurl = resp.json()["destination"]
        if self.bypassedurl == [""]:
            print("[!] Bypass Failed!")
        else:
            print("[+] Bypass Successful! | ", self.bypassedurl)
            print("[+] Bypassed In", self.timetaken, "ms")


if __name__ == '__main__':
    linkbypass().bypass()
