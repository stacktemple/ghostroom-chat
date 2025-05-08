import { Link } from "react-router-dom";

export default function Unauthorized() {
  return (
    <div className="min-h-screen flex items-center justify-center bg-surface text-center px-4">
      <div className="space-y-4">
        <h1 className="text-2xl font-bold text-error">Unauthorized Access</h1>
        <p className="text-text-secondary">
          You are not allowed to access this room. Please join again or return
          home.
        </p>
        <Link
          to="/"
          className="inline-block bg-primary text-text-inverse px-4 py-2 rounded hover:opacity-90 transition"
        >
          ⬅ Go to Home
        </Link>
      </div>
    </div>
  );
}
