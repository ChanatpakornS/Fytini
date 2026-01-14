import { LinkForm } from "@/components/LinkForm";

export default function Home() {
  return (
    <div className="flex min-h-screen items-center justify-center bg-zinc-50 font-sans">
      <main className="flex min-h-screen w-full max-w-3xl flex-col items-center space-y-4 py-32 px-16sm:items-start">
        <h1>Welcome to Fytini! 👋</h1>
        <p>The Simple URL Shortener</p>
        <LinkForm />
      </main>
    </div>
  );
}
