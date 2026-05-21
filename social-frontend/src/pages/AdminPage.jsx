import { useState } from "react";
import { updateUserRole } from "../api/users";

const ROLES = ["user", "teacher", "admin"];

export default function AdminPage() {
  const [userId, setUserId] = useState("");
  const [role, setRole] = useState("teacher");
  const [message, setMessage] = useState("");
  const [error, setError] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();
    setMessage(""); setError("");
    try {
      await updateUserRole(userId, role);
      setMessage(`✅ Роль пользователя #${userId} изменена на "${role}"`);
      setUserId("");
    } catch (err) {
      setError(err.response?.data?.error || "Ошибка");
    }
  };

  return (
    <div style={styles.wrapper}>
      <div style={styles.card}>
        <h2>⚙️ Панель администратора</h2>
        <p style={styles.sub}>Изменение роли пользователя</p>
        <form onSubmit={handleSubmit} style={styles.form}>
          <label style={styles.label}>ID пользователя</label>
          <input
            style={styles.input} type="number" placeholder="Например: 5"
            value={userId} onChange={(e) => setUserId(e.target.value)} required
          />
          <label style={styles.label}>Новая роль</label>
          <select style={styles.input} value={role} onChange={(e) => setRole(e.target.value)}>
            {ROLES.map((r) => <option key={r} value={r}>{r}</option>)}
          </select>
          <button type="submit" style={styles.btn}>Применить</button>
        </form>
        {message && <div style={styles.success}>{message}</div>}
        {error && <div style={styles.error}>{error}</div>}
      </div>
    </div>
  );
}

const styles = {
  wrapper: { display: "flex", justifyContent: "center", padding: "40px 16px" },
  card: { background: "#fff", borderRadius: 12, padding: 32, width: 420, boxShadow: "0 4px 20px rgba(0,0,0,0.1)" },
  sub: { color: "#777", marginBottom: 20 },
  form: { display: "flex", flexDirection: "column", gap: 12 },
  label: { fontWeight: 600, fontSize: 14 },
  input: { padding: "10px 14px", border: "1px solid #ddd", borderRadius: 8, fontSize: 15 },
  btn: { background: "#e94560", color: "#fff", border: "none", borderRadius: 8, padding: "12px", fontSize: 16, cursor: "pointer" },
  success: { marginTop: 12, background: "#efffef", color: "#2a7a2a", padding: 10, borderRadius: 8 },
  error: { marginTop: 12, background: "#fee", color: "#c00", padding: 10, borderRadius: 8 },
};