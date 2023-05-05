import http.client

def main():
    conn = http.client.HTTPConnection("localhost", 8000)

    headers = {
        "User-Agent": "Python-Client",
        "MARK-norgren": "MakdaLLAsl"
    }

    conn.request("GET", "/", headers=headers)

    res = conn.getresponse()
    res_headers = res.getheaders()

    print("Response Headers:")
    for key, value in res_headers:
        print(f"{key}: {value}")

    print("\nResponse Body:")
    print(res.read().decode("utf-8"))

if __name__ == "__main__":
    main()