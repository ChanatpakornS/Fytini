"use client";

import { GetShortenURL } from "@/api/url";
import { usePathname, useRouter } from "next/navigation";
import { useEffect, useState } from "react";

export default function RedirectPage() {
  const pathName = usePathname();
  const router = useRouter();
  const [error, setError] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    async function fetchAndRedirect() {
      // Extract the custom alias from the pathname
      // Remove leading slash and get the alias
      const customAlias = pathName.replace(/^\//, "");

      if (!customAlias) {
        setError("Invalid URL");
        setIsLoading(false);
        return;
      }

      try {
        // Fetch the original URL from the backend
        const result = await GetShortenURL(customAlias);

        if (result.status === "success" && result.url) {
          // Redirect to the original URL
          window.location.href = result.url;
        } else if (result.status === "expired") {
          setError("This URL has expired");
          setIsLoading(false);
        } else if (result.status === "not_found") {
          setError("URL not found");
          setIsLoading(false);
        } else {
          console.error("No URL returned from API or error occurred");
          setError(
            result.message || "An error occurred while fetching the URL",
          );
          setIsLoading(false);
        }
      } catch (err) {
        console.error("Error fetching redirect URL:", err);
        setError("An error occurred while fetching the URL");
        setIsLoading(false);
      }
    }

    fetchAndRedirect();
  }, [pathName, router]);

  if (isLoading) {
    return (
      <div className="flex min-h-screen items-center justify-center bg-zinc-50">
        <div className="text-center space-y-4">
          <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-slate-600 mx-auto"></div>
          <p className="text-lg font-semibold text-slate-700">Redirecting...</p>
          <p className="text-sm text-slate-500">
            Please wait while we fetch your destination
          </p>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="flex min-h-screen items-center justify-center bg-zinc-50">
        <div className="text-center space-y-4 max-w-md">
          <div className="text-6xl">❌</div>
          <h1 className="text-2xl font-bold text-slate-800">Oops!</h1>
          <p className="text-lg text-slate-600">{error}</p>
          <button
            onClick={() => router.push("/")}
            className="mt-4 px-6 py-3 bg-slate-400 text-white rounded-lg font-semibold hover:bg-slate-600/90 duration-300"
          >
            Go back to home
          </button>
        </div>
      </div>
    );
  }

  return null;
}
