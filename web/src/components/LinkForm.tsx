"use client";

import { Button } from "@/components/Button";

function LinkForm() {
  function handleSubmit() {
    alert("Form submitted!");
  }

  return (
    <form className="w-full space-y-4 items-center" onSubmit={handleSubmit}>
      <legend>Full-URL:</legend>
      <input
        type="text"
        placeholder="https://example.com"
        className="border border-gray-300 rounded-md px-4 py-2 w-full"
      />
      <legend>Custom-alias:</legend>
      <input
        type="text"
        placeholder="fytini"
        className="border border-gray-300 rounded-md px-4 py-2 w-full"
      />
      <div className="text-right">
        <Button type="submit" className="w-40">
          Submit
        </Button>
      </div>
    </form>
  );
}

export { LinkForm };
