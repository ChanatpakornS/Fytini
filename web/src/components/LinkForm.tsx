"use client";

import { Button } from "@/components/Button";
import { useState } from "react";

function LinkForm() {
  const [url, setUrl] = useState("");
  const [alias, setAlias] = useState("");
  const [expiration, setExpiration] = useState("");

  function handleSubmit() {
    alert(
      `Form submitted!\n URL: ${url}\n Alias: ${alias}\n Expiration: ${expiration}`,
    );
  }

  return (
    <form className="w-full space-y-4 items-center" onSubmit={handleSubmit}>
      <legend>Full-URL:</legend>
      <input
        type="text"
        placeholder="https://example.com"
        className="border border-gray-300 rounded-md px-4 py-2 w-full"
        onChange={(event) => setUrl(event.target.value)}
      />
      <legend>Custom-alias:</legend>
      <input
        type="text"
        placeholder="fytini"
        className="border border-gray-300 rounded-md px-4 py-2 w-full"
        onChange={(event) => setAlias(event.target.value)}
      />
      <legend>Expiration:</legend>
      <input
        type="datetime-local"
        placeholder="fytini"
        className="border border-gray-300 rounded-md px-4 py-2 w-full"
        onChange={(event) => setExpiration(event.target.value)}
      />
      <div className="text-right mt-4">
        <Button type="submit" className="w-40">
          Submit
        </Button>
      </div>
    </form>
  );
}

export { LinkForm };
