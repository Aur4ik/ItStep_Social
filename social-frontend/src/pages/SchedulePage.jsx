import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import {
  getCommunitySchedule, createSchedule,
  updateSchedule, deleteSchedule,
} from "../api/communities";
import { useAuth } from "../context/AuthContext";

const emptyForm = { title: "", description: "", date: "", time: "" };

export default function SchedulePage() {
  const { id } = useParams();
  const { user } = useAuth();
  const [schedules, setSchedules] = useState([]);
  const [form, setForm] = useState(emptyForm);
  const [editingId, setEditingId] = useState(null);
  const [showForm, setShowForm] = useState(false);
  const [error, setError] = useState("");

  const canManage = user?.role === "teacher" || user?.role === "admin";

  useEffect(() => {
    getCommunitySchedule(id).then((res) => setSchedules(res.data || [])).catch(() => {});
  }, [id]);

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError("");
    try {
      if (editingId) {
        await updateSchedule(editingId, form);
        setSchedules((prev) => prev.map((s) => s.id === editingId ? { ...s, ...form } : s));
        setEditingId(null);
      } else {
        const res = await createSchedule(id, form);
        setSchedules((prev) => [...prev, res.data]);
      }
      setForm(emptyForm);
      setShowForm(false);
    } catch (err) {
      setError(err.response?.data?.error || "Ошибка");
    }
  };

  const handleEdit = (s) => {
    setForm({ title: s.title, description: s.description, date: s.date, time: s.time });
    setEditingId(s.id);
    setShowForm(true);
  };

  const handleDelete = async (scheduleId) => {
    if (!window.confirm("Удалить занятие?")) return;
    await deleteSchedule(scheduleId);
    setSchedules((prev) => prev.filter((s) => s.id !== scheduleId));
  };

  return (
    <div style={styles.wrapper}>
      <div style={styles.header}>
        <h2>📅 Расписание</h2>
        {canManage && (
          <button onClick={() => { setShowForm(!showForm); setEditingId(null); setForm(emptyForm); }} style={styles.btn}>
            {showForm ? "Отмена" : "+ Добавить"}
          </button>
        )}
      </div>

      {showForm && (
        <form onSubmit={handleSubmit} style={styles.form}>
          <h3 style={{ margin: "0 0 12px" }}>{editingId ? "Редактировать" : "Новое занятие"}</h3>
          <input style={styles.input} placeholder="Название" value={form.title} onChange={(e) => setForm({ ...form, title: e.target.value })} required />
          <input style={styles.input} placeholder="Описание" value={form.description} onChange={(e) => setForm({ ...form, description: e.target.value })} />
          <input style={styles.input} type="date" value={form.date} onChange={(e) => setForm({ ...form, date: e.target.value })} required />
          <input style={styles.input} type="time" value={form.time} onChange={(e) => setForm({ ...form, time: e.target.value })} required />
          {error && <div style={styles.error}>{error}</div>}
          <button type="submit" style={styles.btn}>{editingId ? "Сохранить" : "Создать"}</button>
        </form>
      )}

      {schedules.length === 0 ? (
        <div style={styles.empty}>Расписание пока не добавлено</div>
      ) : (
        <div style={styles.list}>
          {schedules.map((s) => (
            <div key={s.id} style={styles.card}>
              <div style={styles.dateBox}>
                <span style={styles.dateText}>{s.date}</span>
                <span style={styles.timeText}>{s.time}</span>
              </div>
              <div style={styles.info}>
                <strong>{s.title}</strong>
                <p style={styles.desc}>{s.description}</p>
                {s.teacher && <span style={styles.teacher}>👤 {s.teacher.first_name} {s.teacher.last_name}</span>}
              </div>
              {canManage && (
                <div style={styles.actions}>
                  <button onClick={() => handleEdit(s)} style={styles.editBtn}>✏️</button>
                  <button onClick={() => handleDelete(s.id)} style={styles.deleteBtn}>🗑️</button>
                </div>
              )}
            </div>
          ))}
        </div>
      )}
    </div>
  );
}

const styles = {
  wrapper: { maxWidth: 700, margin: "24px auto", padding: "0 16px" },
  header: { display: "flex", justifyContent: "space-between", alignItems: "center", marginBottom: 20 },
  btn: { background: "#e94560", color: "#fff", border: "none", borderRadius: 8, padding: "8px 18px", cursor: "pointer" },
  form: { background: "#fff", borderRadius: 10, padding: 20, marginBottom: 20, display: "flex", flexDirection: "column", gap: 10, boxShadow: "0 2px 8px rgba(0,0,0,0.08)" },
  input: { padding: "10px 14px", border: "1px solid #ddd", borderRadius: 8, fontSize: 14 },
  error: { color: "#c00", background: "#fee", padding: 8, borderRadius: 6 },
  empty: { textAlign: "center", color: "#999", marginTop: 40 },
  list: { display: "flex", flexDirection: "column", gap: 12 },
  card: { background: "#fff", borderRadius: 10, padding: 16, display: "flex", alignItems: "center", gap: 16, boxShadow: "0 2px 8px rgba(0,0,0,0.08)" },
  dateBox: { minWidth: 80, background: "#f4f6fb", borderRadius: 8, padding: "8px 12px", textAlign: "center", display: "flex", flexDirection: "column" },
  dateText: { fontSize: 12, color: "#555" },
  timeText: { color: "#e94560", fontWeight: 600, fontSize: 16 },
  info: { flex: 1 },
  desc: { color: "#777", fontSize: 13, margin: "4px 0" },
  teacher: { fontSize: 12, color: "#999" },
  actions: { display: "flex", gap: 8 },
  editBtn: { background: "#f0f0f0", border: "none", borderRadius: 6, padding: "6px 10px", cursor: "pointer" },
  deleteBtn: { background: "#fee", border: "none", borderRadius: 6, padding: "6px 10px", cursor: "pointer" },
};