import { Link, useNavigate } from "react-router-dom";
import { useAuth } from "../context/AuthContext";

export default function Navbar() {
  const { user, logoutUser } = useAuth();
  const navigate = useNavigate();

  const handleLogout = () => {
    logoutUser();
    navigate("/login");
  };

  return (
    <nav style={styles.nav}>
      <Link to="/" style={styles.brand}>📘 ItStep Social</Link>
      <div style={styles.links}>
        {user ? (
          <>
            <Link to="/" style={styles.link}>Лента</Link>
            <Link to="/communities" style={styles.link}>Сообщества</Link>
            <Link to="/chats" style={styles.link}>Чаты</Link>
            <Link to="/profile" style={styles.link}>
              {user.first_name || user.email}
            </Link>
            {user.role === "admin" && (
              <Link to="/admin" style={styles.link}>⚙️ Админ</Link>
            )}
            <button onClick={handleLogout} style={styles.btn}>Выйти</button>
          </>
        ) : (
          <>
            <Link to="/login" style={styles.link}>Войти</Link>
            <Link to="/register" style={styles.link}>Регистрация</Link>
          </>
        )}
      </div>
    </nav>
  );
}

const styles = {
  nav: {
    display: "flex", justifyContent: "space-between", alignItems: "center",
    padding: "12px 24px", background: "#1a1a2e", color: "#fff",
  },
  brand: { color: "#e94560", fontWeight: "bold", fontSize: 20, textDecoration: "none" },
  links: { display: "flex", gap: 16, alignItems: "center" },
  link: { color: "#ccc", textDecoration: "none" },
  btn: {
    background: "#e94560", color: "#fff", border: "none",
    padding: "6px 14px", borderRadius: 6, cursor: "pointer",
  },
};