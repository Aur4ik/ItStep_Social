import { useState } from "react";
import { useNavigate, Link } from "react-router-dom";
import { login } from "../api/auth";
import { useAuth } from "../context/AuthContext";

export default function LoginPage() {
  const [form, setForm] = useState({ email: "", password: "" });
  const [error, setError] = useState("");
  const { loginUser } = useAuth();
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError("");
    try {
      const res = await login(form);
      await loginUser(res.data.token);
      navigate("/");
    } catch (err) {
      setError(err.response?.data?.error || "Ошибка входа");
    }
  };

  return (
    <div style={styles.wrapper}>
      <form onSubmit={handleSubmit} style={styles.form}>
        <h2>Вход</h2>
        {error && <div style={styles.error}>{error}</div>}
        <input
          style={styles.input} placeholder="Email"
          value={form.email} onChange={(e) => setForm({ ...form, email: e.target.value })}
        />
        <input
          style={styles.input} type="password" placeholder="Пароль"
          value={form.password} onChange={(e) => setForm({ ...form, password: e.target.value })}
        />
        <button type="submit" style={styles.btn}>Войти</button>
        <p>Нет аккаунта? <Link to="/register">Регистрация</Link></p>
      </form>
    </div>
  );
}

const styles = {
  wrapper: { display: "flex", justifyContent: "center", alignItems: "center", minHeight: "80vh" },
  form: { background: "#fff", padding: 32, borderRadius: 12, width: 360, boxShadow: "0 4px 20px rgba(0,0,0,0.1)", display: "flex", flexDirection: "column", gap: 12 },
  input: { padding: "10px 14px", border: "1px solid #ddd", borderRadius: 8, fontSize: 15 },
  btn: { background: "#e94560", color: "#fff", border: "none", borderRadius: 8, padding: "12px", fontSize: 16, cursor: "pointer" },
  error: { background: "#fee", color: "#c00", padding: 8, borderRadius: 6 },
};