import { CreateShortenURLRequest, CreateShortenURLResponse } from "@/types/api";

// API configuration
const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8000";
const API_VERSION = "/api/v1";

export type GetUrlResult = {
  url: string | null;
  status: "success" | "expired" | "not_found" | "error";
  message?: string;
};

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
): Promise<GetUrlResult> {
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

      return { url: url, status: "success" };
    }

    // If not found
    if (response.status === 404) {
      console.error("URL not found");
      return { url: null, status: "not_found" };
    }

    // If expired
    if (response.status === 410) {
      console.warn("URL has expired");
      return { url: null, status: "expired" };
    }

    // Try to parse error response
    let errorMessage = response.statusText;
    try {
      const errorData = await response.json();
      errorMessage = errorData.error || errorMessage;
      console.error("Error getting shortened URL:", errorMessage);
    } catch {
      console.error("Error getting shortened URL:", errorMessage);
    }

    return { url: null, status: "error", message: errorMessage };
  } catch (error) {
    console.error("Error fetching redirect URL:", error);
    return { url: null, status: "error", message: "Network error" };
  }
}
