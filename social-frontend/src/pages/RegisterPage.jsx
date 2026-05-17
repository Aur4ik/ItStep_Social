import { useState } from "react";
import { useNavigate, Link } from "react-router-dom";
import { register, login } from "../api/auth";
import { useAuth } from "../context/AuthContext";

export default function RegisterPage() {
  const [form, setForm] = useState({
    email: "", password: "", first_name: "", last_name: "", group: "",
  });
  const [error, setError] = useState("");
  const { loginUser } = useAuth();
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError("");
    try {
      await register(form);
      const res = await login({ email: form.email, password: form.password });
      await loginUser(res.data.token);
      navigate("/");
    } catch (err) {
      setError(err.response?.data?.error || "Ошибка регистрации");
    }
  };

  const field = (key, placeholder, type = "text") => (
    <input
      style={styles.input} type={type} placeholder={placeholder}
      value={form[key]} onChange={(e) => setForm({ ...form, [key]: e.target.value })}
    />
  );

  return (
    <div style={styles.wrapper}>
      <form onSubmit={handleSubmit} style={styles.form}>
        <h2>Регистрация</h2>
        {error && <div style={styles.error}>{error}</div>}
        {field("email", "Email", "email")}
        {field("password", "Пароль", "password")}
        {field("first_name", "Имя")}
        {field("last_name", "Фамилия")}
        {field("group", "Группа")}
        <button type="submit" style={styles.btn}>Зарегистрироваться</button>
        <p>Уже есть аккаунт? <Link to="/login">Войти</Link></p>
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