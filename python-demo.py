import http.client

def main():
    conn = http.client.HTTPSConnection("httpbin.org")

    headers = {
        "Custom-Header": "MyCustomValue",
        "Another-Header": "AnotherValue"
    }

    conn.request("GET", "/response-headers?Custom-Header=MyCustomValue&Another-Header=AnotherValue", headers=headers)

    res = conn.getresponse()
    res_headers = res.getheaders()

    print("Response Headers:")
    for key, value in res_headers:
        print(f"{key}: {value}")

    print("\nResponse Body:")
    print(res.read().decode("utf-8"))

if __name__ == "__main__":
    main()