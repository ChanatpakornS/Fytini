"use client";

import { usePathname } from "next/navigation";
import { useEffect } from "react";

export default function RedirectPage() {
  const pathName = usePathname();

  useEffect(() => {
    alert(pathName);
  }, [pathName]);

  return (
    <div>
      <p>Redirecting...</p>
    </div>
  );
}
