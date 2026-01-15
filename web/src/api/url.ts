import { CreateShortenURLRequest, CreateShortenURLResponse } from "@/types/api";

// API configuration
const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8000";
const API_VERSION = "/api/v1";

// API Functions
export async function CreateNewShortenURL(
  data: CreateShortenURLRequest,
): Promise<CreateShortenURLResponse> {
  try {
    const response = await fetch(
      `${API_BASE_URL}${API_VERSION}/fyt/url/shorten`,
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      },
    );

    if (response.ok) {
      return { message: "URL shortened successfully" };
    }

    // Try to parse error response
    try {
      const errorData = await response.json();
      return { error: errorData.error || "Failed to shorten URL" };
    } catch {
      return { error: `Failed to shorten URL: ${response.statusText}` };
    }
  } catch (error) {
    console.error("Error creating shortened URL:", error);
    return { error: "Network error: Unable to connect to the server" };
  }
}

export async function GetShortenURL(
  custom_alias: string,
): Promise<string | null> {
  try {
    const response = await fetch(
      `${API_BASE_URL}${API_VERSION}/tini/url/redirect`,
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          custom_alias: custom_alias,
        }),
      },
    );

    // Handle success response
    if (response.ok) {
      const data = await response.json();
      const url = data.url;

      console.log("Received URL from API:", url);

      // Return the URL as-is (already normalized when created)
      return url || null;
    }

    // If not found or error
    if (response.status === 404) {
      console.error("URL not found");
      return null;
    }

    // Try to parse error response
    try {
      const errorData = await response.json();
      console.error("Error getting shortened URL:", errorData.error);
    } catch {
      console.error("Error getting shortened URL:", response.statusText);
    }

    return null;
  } catch (error) {
    console.error("Error fetching redirect URL:", error);
    return null;
  }
}
