import Foundation

@available(macOS 12, iOS 15, *)
func fetchData() async {
    let url = URL(string: "http://localhost:8000")!
    var request = URLRequest(url: url)
    request.httpMethod = "GET"
    request.addValue("Swift-Client", forHTTPHeaderField: "User-Agent")
    request.addValue("MyCustomValue", forHTTPHeaderField: "Custom-Header")
    request.addValue("AnotherValue", forHTTPHeaderField: "Another-Header")
    request.addValue("MakdaLLAsl", forHTTPHeaderField: "MARK-norgren")

    do {
        let (data, response) = try await URLSession.shared.data(for: request)
        if let httpResponse = response as? HTTPURLResponse {
            print("Response Headers:")
            for (key, value) in httpResponse.allHeaderFields {
                print("\(key): \(value)")
            }

            print("\nResponse Body:")
            print(String(data: data, encoding: .utf8) ?? "")
        }
    } catch {
        print("Error: \(error)")
    }
}

if #available(macOS 12, iOS 15, *) {
    Task.init {
        await fetchData()
    }
} else {
    // Fallback on earlier versions
    print("Swift concurrency features are not available.")
}

RunLoop.main.run(until: Date(timeIntervalSinceNow: 5))
