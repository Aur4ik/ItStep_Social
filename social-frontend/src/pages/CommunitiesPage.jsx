import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { getCommunities, joinCommunity, createCommunity } from "../api/communities";
import { useAuth } from "../context/AuthContext";

export default function CommunitiesPage() {
  const [communities, setCommunities] = useState([]);
  const [form, setForm] = useState({ name: "", description: "" });
  const [showForm, setShowForm] = useState(false);
  const [error, setError] = useState("");
  const { user } = useAuth();
  const navigate = useNavigate();

  const canManage = user?.role === "teacher" || user?.role === "admin";

  useEffect(() => {
    if (canManage) {
      getCommunities()
        .then((res) => setCommunities(res.data || []))
        .catch(() => setError("Нет доступа к списку сообществ"));
    }
  }, [canManage]);

  const handleJoin = async (id) => {
    await joinCommunity(id);
    navigate(`/communities/${id}`);
  };

  const handleCreate = async (e) => {
    e.preventDefault();
    try {
      const res = await createCommunity(form);
      setCommunities((prev) => [...prev, res.data]);
      setForm({ name: "", description: "" });
      setShowForm(false);
    } catch (err) {
      setError(err.response?.data?.error || "Ошибка создания");
    }
  };

  return (
    <div style={styles.wrapper}>
      <div style={styles.header}>
        <h2>Сообщества</h2>
        {canManage && (
          <button onClick={() => setShowForm(!showForm)} style={styles.btn}>
            {showForm ? "Отмена" : "+ Создать"}
          </button>
        )}
      </div>

      {error && <div style={styles.error}>{error}</div>}

      {showForm && (
        <form onSubmit={handleCreate} style={styles.form}>
          <input
            style={styles.input} placeholder="Название"
            value={form.name} onChange={(e) => setForm({ ...form, name: e.target.value })}
          />
          <input
            style={styles.input} placeholder="Описание"
            value={form.description} onChange={(e) => setForm({ ...form, description: e.target.value })}
          />
          <button type="submit" style={styles.btn}>Создать</button>
        </form>
      )}

      {!canManage && (
        <div style={styles.empty}>
          Список сообществ доступен только преподавателям.<br />
          Попросите преподавателя добавить вас в сообщество.
        </div>
      )}

      <div style={styles.grid}>
        {communities.map((c) => (
          <div key={c.id} style={styles.card}>
            <h3>{c.name}</h3>
            <p style={styles.desc}>{c.description}</p>
            <p style={styles.members}>👥 {c.members_count || 0} участников</p>
            <div style={styles.cardActions}>
              <button onClick={() => navigate(`/communities/${c.id}`)} style={styles.viewBtn}>
                Открыть
              </button>
              <button onClick={() => handleJoin(c.id)} style={styles.joinBtn}>
                Вступить
              </button>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}

const styles = {
  wrapper: { maxWidth: 800, margin: "24px auto", padding: "0 16px" },
  header: { display: "flex", justifyContent: "space-between", alignItems: "center", marginBottom: 20 },
  btn: { background: "#e94560", color: "#fff", border: "none", borderRadius: 8, padding: "8px 18px", cursor: "pointer" },
  form: { background: "#fff", padding: 16, borderRadius: 10, marginBottom: 20, display: "flex", gap: 10, boxShadow: "0 2px 8px rgba(0,0,0,0.08)" },
  input: { flex: 1, padding: "8px 12px", border: "1px solid #ddd", borderRadius: 8 },
  error: { color: "#c00", background: "#fee", padding: 10, borderRadius: 8, marginBottom: 12 },
  empty: { textAlign: "center", color: "#999", marginTop: 40, lineHeight: 1.8 },
  grid: { display: "grid", gridTemplateColumns: "repeat(auto-fill, minmax(220px, 1fr))", gap: 16 },
  card: { background: "#fff", borderRadius: 10, padding: 16, boxShadow: "0 2px 8px rgba(0,0,0,0.08)" },
  desc: { color: "#777", fontSize: 14, margin: "6px 0 4px" },
  members: { color: "#999", fontSize: 12, margin: "0 0 12px" },
  cardActions: { display: "flex", gap: 8 },
  viewBtn: { flex: 1, background: "#f0f0f0", border: "none", borderRadius: 6, padding: "6px", cursor: "pointer" },
  joinBtn: { flex: 1, background: "#e94560", color: "#fff", border: "none", borderRadius: 6, padding: "6px", cursor: "pointer" },
};