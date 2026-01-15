"use client";

import { CreateNewShortenURL } from "@/api/url";
import { Button } from "@/components/Button";
import { formatExpirationDate } from "@/utils/date";
import { isValidURL } from "@/utils/url";
import { useState } from "react";

function LinkForm() {
  const [url, setUrl] = useState("");
  const [alias, setAlias] = useState("");
  const [expiration, setExpiration] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const [message, setMessage] = useState<{
    type: "success" | "error";
    text: string;
  } | null>(null);

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    setMessage(null);

    // Validation
    if (!url || !alias) {
      setMessage({
        type: "error",
        text: "URL and Custom Alias are required",
      });
      return;
    }

    // Normalize URL - add protocol if missing
    let normalizedUrl = url.trim();
    if (!normalizedUrl.match(/^https?:\/\//i)) {
      normalizedUrl = "https://" + normalizedUrl;
    }

    if (!isValidURL(normalizedUrl)) {
      setMessage({
        type: "error",
        text: "Please enter a valid URL (e.g., https://example.com)",
      });
      return;
    }

    if (!/^[a-zA-Z0-9_-]+$/.test(alias)) {
      setMessage({
        type: "error",
        text: "Custom alias can only contain letters, numbers, hyphens, and underscores",
      });
      return;
    }

    setIsLoading(true);

    try {
      const payload = {
        url: normalizedUrl,
        custom_alias: alias,
        expiration_date: expiration ? formatExpirationDate(expiration) : "",
      };

      const response = await CreateNewShortenURL(payload);

      if (response.error) {
        setMessage({
          type: "error",
          text: response.error,
        });
      } else {
        setMessage({
          type: "success",
          text: `URL shortened successfully! Your short URL: ${window.location.origin}/${alias}`,
        });
        // Reset form
        setUrl("");
        setAlias("");
        setExpiration("");
      }
    } catch {
      setMessage({
        type: "error",
        text: "An unexpected error occurred. Please try again.",
      });
    } finally {
      setIsLoading(false);
    }
  }

  return (
    <form className="w-full space-y-4 items-center" onSubmit={handleSubmit}>
      {message && (
        <div
          className={`p-4 rounded-md ${
            message.type === "success"
              ? "bg-green-100 text-green-800 border border-green-300"
              : "bg-red-100 text-red-800 border border-red-300"
          }`}
        >
          {message.text}
        </div>
      )}

      <legend className="font-semibold">Full-URL:</legend>
      <input
        type="text"
        placeholder="https://example.com"
        className="border border-gray-300 rounded-md px-4 py-2 w-full focus:outline-none focus:ring-2 focus:ring-slate-400"
        value={url}
        onChange={(event) => setUrl(event.target.value)}
        disabled={isLoading}
        required
      />

      <legend className="font-semibold">Custom-alias:</legend>
      <input
        type="text"
        placeholder="fytini"
        className="border border-gray-300 rounded-md px-4 py-2 w-full focus:outline-none focus:ring-2 focus:ring-slate-400"
        value={alias}
        onChange={(event) => setAlias(event.target.value)}
        disabled={isLoading}
        required
      />

      <legend className="font-semibold">
        Expiration <span className="font-normal text-gray-500">(optional)</span>
        :
      </legend>
      <input
        type="datetime-local"
        className="border border-gray-300 rounded-md px-4 py-2 w-full focus:outline-none focus:ring-2 focus:ring-slate-400"
        value={expiration}
        onChange={(event) => setExpiration(event.target.value)}
        disabled={isLoading}
      />

      <div className="text-right mt-4">
        <Button type="submit" className="w-40" disabled={isLoading}>
          {isLoading ? "Submitting..." : "Submit"}
        </Button>
      </div>
    </form>
  );
}

export { LinkForm };
