import { Link, Outlet } from "react-router-dom";
export default function Layout() {
  return (
    <div className="min-h-screen flex flex-col">
      <header className="bg-surface-header text-text-inverse p-4 text-lg font-semibold">
        <Link to="/" className=" text-white no-underline hover:opacity-80">
          Emperor Chat ☯
        </Link>
      </header>
      <main className="flex-1 bg-surface">
        <Outlet />
      </main>
      <footer className="bg-surface-footer text-text-inverse text-center text-sm p-2">
        © 2025 Stacktemple (Personal Project)
      </footer>
    </div>
  );
}
