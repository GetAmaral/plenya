import { useState, useEffect, useRef } from "react";

// ═══════════════════════════════════════════════════════════════
// NORMATIVE DATA FROM UPLOADED DOCUMENTS
// ═══════════════════════════════════════════════════════════════

// ── Abdominal Test (Pollock & Wilmore, 1993) - reps/min ──
const ABDOMINAL_NORMS = {
  M: [
    { age: "15-19", excelente: 48, acima: [42,47], media: [38,41], abaixo: [33,37], fraco: 32 },
    { age: "20-29", excelente: 43, acima: [37,42], media: [33,36], abaixo: [29,32], fraco: 28 },
    { age: "30-39", excelente: 36, acima: [31,35], media: [27,30], abaixo: [22,26], fraco: 21 },
    { age: "40-49", excelente: 31, acima: [26,30], media: [22,25], abaixo: [17,21], fraco: 16 },
    { age: "50-59", excelente: 26, acima: [22,25], media: [18,21], abaixo: [13,17], fraco: 12 },
    { age: "60-69", excelente: 23, acima: [17,22], media: [12,16], abaixo: [7,11], fraco: 6 },
  ],
  F: [
    { age: "15-19", excelente: 42, acima: [36,41], media: [32,35], abaixo: [27,31], fraco: 26 },
    { age: "20-29", excelente: 36, acima: [31,35], media: [25,30], abaixo: [21,24], fraco: 20 },
    { age: "30-39", excelente: 29, acima: [24,28], media: [20,23], abaixo: [15,19], fraco: 14 },
    { age: "40-49", excelente: 25, acima: [20,24], media: [15,19], abaixo: [7,14], fraco: 6 },
    { age: "50-59", excelente: 19, acima: [12,18], media: [5,11], abaixo: [3,4], fraco: 2 },
    { age: "60-69", excelente: 16, acima: [12,15], media: [4,11], abaixo: [2,3], fraco: 1 },
  ]
};

// ── Push-up / Flexão de Solo (Pollock & Wilmore, 1993) - reps/min ──
const PUSHUP_NORMS = {
  M: [
    { age: "15-19", excelente: 39, acima: [29,38], media: [23,28], abaixo: [18,22], fraco: 17 },
    { age: "20-29", excelente: 36, acima: [29,35], media: [22,28], abaixo: [17,21], fraco: 16 },
    { age: "30-39", excelente: 30, acima: [22,29], media: [17,21], abaixo: [12,16], fraco: 11 },
    { age: "40-49", excelente: 22, acima: [17,21], media: [13,16], abaixo: [10,12], fraco: 9 },
    { age: "50-59", excelente: 21, acima: [13,20], media: [10,12], abaixo: [7,9], fraco: 6 },
    { age: "60-69", excelente: 18, acima: [11,17], media: [8,10], abaixo: [5,7], fraco: 4 },
  ],
  F: [
    { age: "15-19", excelente: 33, acima: [25,32], media: [18,24], abaixo: [12,17], fraco: 11 },
    { age: "20-29", excelente: 30, acima: [21,29], media: [15,20], abaixo: [10,14], fraco: 9 },
    { age: "30-39", excelente: 27, acima: [20,26], media: [13,19], abaixo: [8,12], fraco: 7 },
    { age: "40-49", excelente: 24, acima: [15,23], media: [11,14], abaixo: [5,10], fraco: 4 },
    { age: "50-59", excelente: 21, acima: [11,22], media: [7,10], abaixo: [2,6], fraco: 1 },
    { age: "60-69", excelente: 17, acima: [12,16], media: [5,11], abaixo: [2,4], fraco: 1 },
  ]
};

// ── Plank Test Percentiles (Strand et al., JHK 2014) - seconds ──
const PLANK_NORMS = {
  M: { p10:62, p20:79, p30:89, p40:97, p50:110, p60:122, p70:137, p80:157, p90:201 },
  F: { p10:35, p20:48, p30:58, p40:63, p50:72, p60:84, p70:95, p80:108, p90:142 },
  M_athlete: { p10:74, p20:84, p30:94, p40:117, p50:125, p60:140, p70:157, p80:183, p90:228 },
  F_athlete: { p10:45, p20:59, p30:63, p40:74, p50:87, p60:97, p70:110, p80:162, p90:194 },
};

// ── 3-Min Burpee Test (Podstawski et al., JHK 2019) ──
const BURPEE_NORMS = {
  M: { muito_fraco: [28,38], fraco: [38,47], medio: [47,66], bom: [66,76], muito_bom: [76,85] },
  F: { muito_fraco: [15,26], fraco: [26,37], medio: [37,60], bom: [60,72], muito_bom: [72,83] },
};

// ── Standard Push-up Scale for Females 20-29 (Adams et al., IJES 2022) ──
const PUSHUP_STANDARD_FEMALE = {
  excelente: 18, muito_bom: [12,17], bom: [8,11], regular: [5,7], fraco: 4
};

// ═══════════════════════════════════════════════════════════════
// CLASSIFICATION FUNCTIONS
// ═══════════════════════════════════════════════════════════════

function classifyPollock(value, norms, sex, ageGroup) {
  const table = norms[sex];
  const row = table.find(r => r.age === ageGroup);
  if (!row) return { class: "N/A", color: "#64748b", score: 0 };
  if (value >= row.excelente) return { class: "Excelente", color: "#10b981", score: 5, emoji: "🏆" };
  if (value >= row.acima[0] && value <= row.acima[1]) return { class: "Acima da Média", color: "#22d3ee", score: 4, emoji: "💪" };
  if (value >= row.media[0] && value <= row.media[1]) return { class: "Média", color: "#eab308", score: 3, emoji: "👍" };
  if (value >= row.abaixo[0] && value <= row.abaixo[1]) return { class: "Abaixo da Média", color: "#f97316", score: 2, emoji: "⚠️" };
  return { class: "Fraco", color: "#ef4444", score: 1, emoji: "🔴" };
}

function classifyPlank(seconds, sex, isAthlete) {
  const key = isAthlete ? `${sex}_athlete` : sex;
  const n = PLANK_NORMS[key] || PLANK_NORMS[sex];
  if (seconds >= n.p90) return { class: "Excelente", color: "#10b981", percentile: "90+", score: 5, emoji: "🏆" };
  if (seconds >= n.p80) return { class: "Muito Bom", color: "#22d3ee", percentile: "80-89", score: 4, emoji: "💪" };
  if (seconds >= n.p60) return { class: "Bom", color: "#a3e635", percentile: "60-79", score: 4, emoji: "👍" };
  if (seconds >= n.p40) return { class: "Médio", color: "#eab308", percentile: "40-59", score: 3, emoji: "👌" };
  if (seconds >= n.p20) return { class: "Abaixo da Média", color: "#f97316", percentile: "20-39", score: 2, emoji: "⚠️" };
  return { class: "Fraco", color: "#ef4444", percentile: "<20", score: 1, emoji: "🔴" };
}

function classifyBurpee(cycles, sex) {
  const n = BURPEE_NORMS[sex];
  if (cycles >= n.muito_bom[0]) return { class: "Muito Bom", color: "#10b981", score: 5, emoji: "🏆", tScale: Math.min(100, 81 + (cycles - n.muito_bom[0]) * 2) };
  if (cycles >= n.bom[0]) return { class: "Bom", color: "#22d3ee", score: 4, emoji: "💪", tScale: 61 + Math.round((cycles - n.bom[0]) / (n.bom[1] - n.bom[0]) * 19) };
  if (cycles >= n.medio[0]) return { class: "Médio", color: "#eab308", score: 3, emoji: "👍", tScale: 41 + Math.round((cycles - n.medio[0]) / (n.medio[1] - n.medio[0]) * 19) };
  if (cycles >= n.fraco[0]) return { class: "Fraco", color: "#f97316", score: 2, emoji: "⚠️", tScale: 21 + Math.round((cycles - n.fraco[0]) / (n.fraco[1] - n.fraco[0]) * 19) };
  return { class: "Muito Fraco", color: "#ef4444", score: 1, emoji: "🔴", tScale: Math.max(1, Math.round((cycles / n.muito_fraco[1]) * 20)) };
}

const AGE_GROUPS = ["15-19","20-29","30-39","40-49","50-59","60-69"];

function getAgeGroup(age) {
  if (age < 20) return "15-19";
  if (age < 30) return "20-29";
  if (age < 40) return "30-39";
  if (age < 50) return "40-49";
  if (age < 60) return "50-59";
  return "60-69";
}

// ═══════════════════════════════════════════════════════════════
// MAIN APP
// ═══════════════════════════════════════════════════════════════

export default function FitnessEvalApp() {
  const [screen, setScreen] = useState("home");
  const [profile, setProfile] = useState({ name: "", age: 25, sex: "M", isAthlete: false, weight: "", height: "" });
  const [results, setResults] = useState({ abdominal: null, pushup: null, plank: null, burpee: null });
  const [inputVal, setInputVal] = useState("");
  const [currentTest, setCurrentTest] = useState(null);
  const [evaluations, setEvaluations] = useState([]);
  const [timer, setTimer] = useState(0);
  const [timerRunning, setTimerRunning] = useState(false);
  const [timerMode, setTimerMode] = useState(null); // 'stopwatch' or 'countdown'
  const timerRef = useRef(null);

  // Timer logic
  useEffect(() => {
    if (timerRunning) {
      timerRef.current = setInterval(() => {
        setTimer(prev => {
          if (timerMode === "countdown") {
            if (prev <= 1) { setTimerRunning(false); return 0; }
            return prev - 1;
          }
          return prev + 1;
        });
      }, 1000);
    }
    return () => clearInterval(timerRef.current);
  }, [timerRunning, timerMode]);

  const formatTime = (s) => `${Math.floor(s / 60).toString().padStart(2, "0")}:${(s % 60).toString().padStart(2, "0")}`;

  const ageGroup = getAgeGroup(profile.age);

  const startTest = (testId) => {
    setCurrentTest(testId);
    setInputVal("");
    setTimer(0);
    setTimerRunning(false);
    setScreen("test");
  };

  const saveResult = (testId, value) => {
    const numVal = parseFloat(value);
    if (isNaN(numVal)) return;
    let classification;
    if (testId === "abdominal") classification = classifyPollock(numVal, ABDOMINAL_NORMS, profile.sex, ageGroup);
    else if (testId === "pushup") classification = classifyPollock(numVal, PUSHUP_NORMS, profile.sex, ageGroup);
    else if (testId === "plank") classification = classifyPlank(numVal, profile.sex, profile.isAthlete);
    else if (testId === "burpee") classification = classifyBurpee(numVal, profile.sex);
    setResults(prev => ({ ...prev, [testId]: { value: numVal, ...classification } }));
    setScreen("tests");
  };

  const finishEvaluation = () => {
    const completed = Object.entries(results).filter(([, v]) => v !== null);
    if (completed.length === 0) return;
    const avgScore = completed.reduce((sum, [, v]) => sum + v.score, 0) / completed.length;
    const evaluation = {
      date: new Date().toLocaleString("pt-BR"),
      profile: { ...profile },
      results: { ...results },
      avgScore,
      overallClass: avgScore >= 4.5 ? "Excelente" : avgScore >= 3.5 ? "Bom" : avgScore >= 2.5 ? "Médio" : avgScore >= 1.5 ? "Abaixo da Média" : "Fraco",
    };
    setEvaluations(prev => [evaluation, ...prev]);
    setScreen("report");
  };

  const overallColor = (s) => s >= 4.5 ? "#10b981" : s >= 3.5 ? "#22d3ee" : s >= 2.5 ? "#eab308" : s >= 1.5 ? "#f97316" : "#ef4444";

  const TESTS = [
    { id: "abdominal", name: "Teste Abdominal", icon: "🏋️", unit: "reps/min", desc: "Repetições em 1 minuto (Pollock & Wilmore, 1993)", timerPreset: 60, protocol: "Posição supina, joelhos flexionados a 90°. Mãos cruzadas sobre o peito. Realizar o máximo de abdominais completos em 1 minuto." },
    { id: "pushup", name: "Flexão de Solo", icon: "💪", unit: "reps/min", desc: "Repetições em 1 minuto (Pollock & Wilmore, 1993)", timerPreset: 60, protocol: "Posição de prancha com braços estendidos. Flexionar cotovelos até 90°. Realizar o máximo de flexões em 1 minuto. Mulheres: flexão modificada (joelhos)." },
    { id: "plank", name: "Teste de Prancha", icon: "🧘", unit: "segundos", desc: "Tempo máximo de sustentação (Strand et al., JHK 2014)", timerPreset: 0, protocol: "Prancha frontal sobre antebraços. Cotovelos sob os ombros. Corpo alinhado (sem elevar quadril). Manter até a falha técnica ou volitiva." },
    { id: "burpee", name: "Burpee 3 Min", icon: "🔥", unit: "ciclos/3min", desc: "Ciclos completos em 3 minutos (Podstawski et al., JHK 2019)", timerPreset: 180, protocol: "De pé → agachar com mãos no chão → chutar pés para trás (prancha) → retornar ao agachamento → ficar de pé com mãos acima da cabeça. Repetir o máximo possível em 3 minutos." },
  ];

  // ─── HOME ──────────────────────────────────────────────────
  if (screen === "home") {
    return (
      <div style={S.app}>
        <div style={S.homeWrap}>
          <div style={S.brand}>
            <div style={S.brandIcon}>
              <svg viewBox="0 0 48 48" width="56" height="56">
                <rect x="4" y="20" width="8" height="24" rx="2" fill="#10b981" opacity="0.7"/>
                <rect x="14" y="12" width="8" height="32" rx="2" fill="#22d3ee" opacity="0.8"/>
                <rect x="24" y="6" width="8" height="38" rx="2" fill="#6366f1"/>
                <rect x="34" y="14" width="8" height="30" rx="2" fill="#f59e0b" opacity="0.8"/>
              </svg>
            </div>
            <h1 style={S.brandTitle}>FitEval</h1>
            <p style={S.brandSub}>Sistema de Avaliação Física</p>
            <p style={S.brandRef}>Baseado em evidências científicas</p>
          </div>

          <div style={S.profileCard}>
            <h3 style={S.sectionTitle}>👤 Perfil do Avaliado</h3>
            <div style={S.formGrid}>
              <div style={S.field}>
                <label style={S.label}>Nome</label>
                <input style={S.input} placeholder="Nome completo" value={profile.name} onChange={e => setProfile(p => ({...p, name: e.target.value}))}/>
              </div>
              <div style={S.fieldRow}>
                <div style={{flex:1}}>
                  <label style={S.label}>Idade</label>
                  <input style={S.input} type="number" value={profile.age} onChange={e => setProfile(p => ({...p, age: parseInt(e.target.value)||20}))}/>
                </div>
                <div style={{flex:1}}>
                  <label style={S.label}>Sexo</label>
                  <div style={S.segmented}>
                    {[["M","Masc."],["F","Fem."]].map(([v,l]) => (
                      <button key={v} style={{...S.segBtn, ...(profile.sex===v?S.segBtnActive:{})}} onClick={() => setProfile(p => ({...p, sex: v}))}>{l}</button>
                    ))}
                  </div>
                </div>
              </div>
              <div style={S.fieldRow}>
                <div style={{flex:1}}>
                  <label style={S.label}>Peso (kg)</label>
                  <input style={S.input} type="number" placeholder="70" value={profile.weight} onChange={e => setProfile(p => ({...p, weight: e.target.value}))}/>
                </div>
                <div style={{flex:1}}>
                  <label style={S.label}>Altura (cm)</label>
                  <input style={S.input} type="number" placeholder="170" value={profile.height} onChange={e => setProfile(p => ({...p, height: e.target.value}))}/>
                </div>
              </div>
              <div style={S.checkRow}>
                <label style={{...S.checkLabel, cursor:"pointer"}} onClick={() => setProfile(p => ({...p, isAthlete: !p.isAthlete}))}>
                  <span style={{...S.checkbox, ...(profile.isAthlete ? S.checkboxActive : {})}}>
                    {profile.isAthlete && "✓"}
                  </span>
                  Atleta (Varsity/Competitivo)
                </label>
              </div>
            </div>
          </div>

          <button style={S.primaryBtn} onClick={() => setScreen("tests")}>
            <span style={{fontSize:20}}>📋</span> Iniciar Avaliação
          </button>

          {evaluations.length > 0 && (
            <button style={S.ghostBtn} onClick={() => setScreen("history")}>
              📊 Histórico ({evaluations.length})
            </button>
          )}

          <div style={S.refs}>
            <p style={S.refTitle}>Referências Científicas:</p>
            <p style={S.refText}>• Pollock & Wilmore (1993) — Abdominal e Flexão</p>
            <p style={S.refText}>• Strand et al. (JHK, 2014) — Teste de Prancha</p>
            <p style={S.refText}>• Podstawski et al. (JHK, 2019) — Burpee 3 Min</p>
            <p style={S.refText}>• Adams et al. (IJES, 2022) — Push-up Standard</p>
            <p style={S.refText}>• Brotons-Gil et al. (JSCR, 2013) — FRT Test</p>
          </div>
        </div>
      </div>
    );
  }

  // ─── TEST LIST ─────────────────────────────────────────────
  if (screen === "tests") {
    const completedCount = Object.values(results).filter(r => r !== null).length;
    return (
      <div style={S.app}>
        <div style={S.header}>
          <button style={S.backBtn} onClick={() => setScreen("home")}>← Perfil</button>
          <h2 style={S.headerTitle}>Testes</h2>
          <span style={S.badge}>{completedCount}/{TESTS.length}</span>
        </div>
        <div style={S.testListWrap}>
          <div style={S.profileMini}>
            <span>{profile.name || "Avaliado"}</span>
            <span style={S.profileTag}>{profile.sex === "M" ? "♂" : "♀"} {profile.age}a • {ageGroup}</span>
          </div>

          {TESTS.map(test => {
            const result = results[test.id];
            const done = result !== null;
            return (
              <div key={test.id} style={{...S.testCard, ...(done ? {borderColor: result.color} : {})}}>
                <div style={S.testCardTop}>
                  <span style={{fontSize:28}}>{test.icon}</span>
                  <div style={{flex:1}}>
                    <p style={S.testName}>{test.name}</p>
                    <p style={S.testDesc}>{test.desc}</p>
                  </div>
                  {done ? (
                    <div style={{textAlign:"right"}}>
                      <span style={{fontSize:20,fontWeight:800,color:result.color}}>{result.value}</span>
                      <span style={{display:"block",fontSize:10,color:result.color,fontWeight:600}}>{result.class}</span>
                    </div>
                  ) : (
                    <button style={S.startBtn} onClick={() => startTest(test.id)}>Iniciar</button>
                  )}
                </div>
                {done && (
                  <div style={{marginTop:8,display:"flex",gap:8,alignItems:"center"}}>
                    <div style={{flex:1,height:4,background:"rgba(100,116,139,0.15)",borderRadius:2,overflow:"hidden"}}>
                      <div style={{height:"100%",width:`${result.score*20}%`,background:result.color,borderRadius:2}}/>
                    </div>
                    <span style={{fontSize:11,color:"#94a3b8"}}>{result.emoji} {result.score}/5</span>
                    <button style={S.redoBtn} onClick={() => startTest(test.id)}>↺</button>
                  </div>
                )}
              </div>
            );
          })}

          {completedCount > 0 && (
            <button style={{...S.primaryBtn, marginTop:8}} onClick={finishEvaluation}>
              📊 Gerar Relatório
            </button>
          )}
        </div>
      </div>
    );
  }

  // ─── INDIVIDUAL TEST ───────────────────────────────────────
  if (screen === "test" && currentTest) {
    const test = TESTS.find(t => t.id === currentTest);
    const isPlank = currentTest === "plank";
    return (
      <div style={S.app}>
        <div style={S.header}>
          <button style={S.backBtn} onClick={() => {setTimerRunning(false);setScreen("tests")}}>← Voltar</button>
          <h2 style={S.headerTitle}>{test.icon} {test.name}</h2>
          <div style={{width:50}}/>
        </div>
        <div style={{padding:"16px 20px"}}>
          <div style={S.protocolCard}>
            <p style={{fontSize:13,fontWeight:700,color:"#e2e8f0",margin:"0 0 6px"}}>📋 Protocolo</p>
            <p style={{fontSize:12,color:"#c7d2fe",margin:0,lineHeight:1.6}}>{test.protocol}</p>
          </div>

          {/* Timer */}
          <div style={S.timerCard}>
            <div style={S.timerDisplay}>{formatTime(timer)}</div>
            <div style={{display:"flex",gap:8,justifyContent:"center",flexWrap:"wrap"}}>
              {!timerRunning ? (
                <>
                  {isPlank ? (
                    <button style={S.timerBtn} onClick={() => {setTimerMode("stopwatch");setTimer(0);setTimerRunning(true)}}>
                      ▶ Cronômetro
                    </button>
                  ) : (
                    <>
                      <button style={S.timerBtn} onClick={() => {setTimerMode("countdown");setTimer(test.timerPreset);setTimerRunning(true)}}>
                        ▶ {test.timerPreset}s
                      </button>
                      <button style={{...S.timerBtn,background:"rgba(34,211,238,0.15)",borderColor:"#22d3ee",color:"#22d3ee"}} onClick={() => {setTimerMode("stopwatch");setTimer(0);setTimerRunning(true)}}>
                        ⏱ Livre
                      </button>
                    </>
                  )}
                </>
              ) : (
                <button style={{...S.timerBtn,background:"rgba(239,68,68,0.15)",borderColor:"#ef4444",color:"#ef4444"}} onClick={() => {setTimerRunning(false);if(isPlank){setInputVal(String(timer))}}}>
                  ⏹ Parar {isPlank && `(${timer}s)`}
                </button>
              )}
              {!timerRunning && timer > 0 && (
                <button style={S.timerBtnSmall} onClick={() => setTimer(0)}>Zerar</button>
              )}
            </div>
          </div>

          {/* Input */}
          <div style={S.inputCard}>
            <label style={S.label}>Resultado ({test.unit})</label>
            <div style={{display:"flex",gap:10,alignItems:"center"}}>
              <input style={{...S.input, flex:1, fontSize:24, textAlign:"center", fontWeight:700}}
                type="number" placeholder="0" value={inputVal}
                onChange={e => setInputVal(e.target.value)}/>
              <button style={{...S.primaryBtn, padding:"14px 24px"}} onClick={() => saveResult(currentTest, inputVal)} disabled={!inputVal}>
                Salvar
              </button>
            </div>
          </div>

          {/* Normative Table Preview */}
          {(currentTest === "abdominal" || currentTest === "pushup") && (
            <div style={S.tableCard}>
              <p style={{fontSize:12,fontWeight:700,color:"#e2e8f0",margin:"0 0 8px"}}>
                📊 Tabela Normativa — {profile.sex === "M" ? "Homens" : "Mulheres"} ({ageGroup} anos)
              </p>
              {(() => {
                const norms = currentTest === "abdominal" ? ABDOMINAL_NORMS : PUSHUP_NORMS;
                const row = norms[profile.sex].find(r => r.age === ageGroup);
                if (!row) return null;
                const cats = [
                  { name: "Excelente", val: `≥ ${row.excelente}`, color: "#10b981" },
                  { name: "Acima Média", val: `${row.acima[0]}-${row.acima[1]}`, color: "#22d3ee" },
                  { name: "Média", val: `${row.media[0]}-${row.media[1]}`, color: "#eab308" },
                  { name: "Abaixo Média", val: `${row.abaixo[0]}-${row.abaixo[1]}`, color: "#f97316" },
                  { name: "Fraco", val: `≤ ${row.fraco}`, color: "#ef4444" },
                ];
                return cats.map((c, i) => (
                  <div key={i} style={{display:"flex",justifyContent:"space-between",padding:"5px 0",borderBottom: i<4 ? "1px solid rgba(100,116,139,0.1)":"none"}}>
                    <span style={{fontSize:12,color:c.color,fontWeight:600}}>{c.name}</span>
                    <span style={{fontSize:12,color:"#e2e8f0",fontWeight:700}}>{c.val}</span>
                  </div>
                ));
              })()}
            </div>
          )}

          {currentTest === "plank" && (
            <div style={S.tableCard}>
              <p style={{fontSize:12,fontWeight:700,color:"#e2e8f0",margin:"0 0 8px"}}>
                📊 Percentis — {profile.sex === "M" ? "Homens" : "Mulheres"} {profile.isAthlete ? "(Atletas)" : ""}
              </p>
              {(() => {
                const key = profile.isAthlete ? `${profile.sex}_athlete` : profile.sex;
                const n = PLANK_NORMS[key] || PLANK_NORMS[profile.sex];
                return [["P90+","Excelente",n.p90,"#10b981"],["P80","Muito Bom",n.p80,"#22d3ee"],["P60","Bom",n.p60,"#a3e635"],["P40","Médio",n.p40,"#eab308"],["P20","Abaixo",n.p20,"#f97316"],["<P20","Fraco","<"+n.p20,"#ef4444"]].map(([pct,lbl,val,col],i) => (
                  <div key={i} style={{display:"flex",justifyContent:"space-between",padding:"5px 0",borderBottom:i<5?"1px solid rgba(100,116,139,0.1)":"none"}}>
                    <span style={{fontSize:11,color:"#94a3b8"}}>{pct}</span>
                    <span style={{fontSize:12,color:col,fontWeight:600}}>{lbl}</span>
                    <span style={{fontSize:12,color:"#e2e8f0",fontWeight:700}}>{typeof val === "number" ? `≥${val}s` : val+"s"}</span>
                  </div>
                ));
              })()}
            </div>
          )}

          {currentTest === "burpee" && (
            <div style={S.tableCard}>
              <p style={{fontSize:12,fontWeight:700,color:"#e2e8f0",margin:"0 0 8px"}}>
                📊 Standards — {profile.sex === "M" ? "Homens" : "Mulheres"} 18-25 anos
              </p>
              {(() => {
                const n = BURPEE_NORMS[profile.sex];
                return [
                  ["Muito Bom",`${n.muito_bom[0]}-${n.muito_bom[1]}`,"#10b981"],
                  ["Bom",`${n.bom[0]}-${n.bom[1]}`,"#22d3ee"],
                  ["Médio",`${n.medio[0]}-${n.medio[1]}`,"#eab308"],
                  ["Fraco",`${n.fraco[0]}-${n.fraco[1]}`,"#f97316"],
                  ["Muito Fraco",`${n.muito_fraco[0]}-${n.muito_fraco[1]}`,"#ef4444"],
                ].map(([lbl,val,col],i) => (
                  <div key={i} style={{display:"flex",justifyContent:"space-between",padding:"5px 0",borderBottom:i<4?"1px solid rgba(100,116,139,0.1)":"none"}}>
                    <span style={{fontSize:12,color:col,fontWeight:600}}>{lbl}</span>
                    <span style={{fontSize:12,color:"#e2e8f0",fontWeight:700}}>{val} ciclos</span>
                  </div>
                ));
              })()}
            </div>
          )}
        </div>
      </div>
    );
  }

  // ─── REPORT ────────────────────────────────────────────────
  if (screen === "report") {
    const eval0 = evaluations[0];
    if (!eval0) { setScreen("home"); return null; }
    const completed = Object.entries(eval0.results).filter(([, v]) => v !== null);
    const bmi = eval0.profile.weight && eval0.profile.height ? (parseFloat(eval0.profile.weight) / ((parseFloat(eval0.profile.height)/100)**2)).toFixed(1) : null;

    return (
      <div style={S.app}>
        <div style={S.header}>
          <button style={S.backBtn} onClick={() => setScreen("tests")}>← Testes</button>
          <h2 style={S.headerTitle}>Relatório</h2>
          <button style={S.backBtn} onClick={() => setScreen("home")}>🏠</button>
        </div>

        <div style={{padding:"16px 20px"}}>
          {/* Header */}
          <div style={S.reportHeader}>
            <div style={S.reportLogo}>
              <svg viewBox="0 0 32 32" width="32" height="32"><rect x="2" y="14" width="6" height="16" rx="1" fill="#10b981"/><rect x="9" y="8" width="6" height="22" rx="1" fill="#6366f1"/><rect x="16" y="4" width="6" height="26" rx="1" fill="#22d3ee"/><rect x="23" y="10" width="6" height="20" rx="1" fill="#f59e0b"/></svg>
            </div>
            <div>
              <h3 style={{margin:0,fontSize:16,fontWeight:800,color:"#e2e8f0"}}>Relatório de Avaliação Física</h3>
              <p style={{margin:"2px 0 0",fontSize:11,color:"#94a3b8"}}>{eval0.date}</p>
            </div>
          </div>

          {/* Profile Summary */}
          <div style={S.reportProfile}>
            <div style={S.rpRow}><span style={S.rpLabel}>Nome:</span><span style={S.rpVal}>{eval0.profile.name || "—"}</span></div>
            <div style={S.rpRow}><span style={S.rpLabel}>Idade:</span><span style={S.rpVal}>{eval0.profile.age} anos ({ageGroup})</span></div>
            <div style={S.rpRow}><span style={S.rpLabel}>Sexo:</span><span style={S.rpVal}>{eval0.profile.sex === "M" ? "Masculino" : "Feminino"}</span></div>
            {eval0.profile.weight && <div style={S.rpRow}><span style={S.rpLabel}>Peso:</span><span style={S.rpVal}>{eval0.profile.weight} kg</span></div>}
            {eval0.profile.height && <div style={S.rpRow}><span style={S.rpLabel}>Altura:</span><span style={S.rpVal}>{eval0.profile.height} cm</span></div>}
            {bmi && <div style={S.rpRow}><span style={S.rpLabel}>IMC:</span><span style={S.rpVal}>{bmi} kg/m²</span></div>}
            {eval0.profile.isAthlete && <div style={S.rpRow}><span style={S.rpLabel}>Status:</span><span style={{...S.rpVal,color:"#22d3ee"}}>Atleta</span></div>}
          </div>

          {/* Overall Score */}
          <div style={{...S.scoreCard, borderColor: overallColor(eval0.avgScore)}}>
            <div style={{fontSize:52,fontWeight:900,color:overallColor(eval0.avgScore),lineHeight:1}}>
              {eval0.avgScore.toFixed(1)}
            </div>
            <div style={{fontSize:14,fontWeight:700,color:"#c7d2fe",marginTop:4}}>{eval0.overallClass}</div>
            <div style={{height:6,background:"rgba(100,116,139,0.2)",borderRadius:3,marginTop:12,overflow:"hidden"}}>
              <div style={{height:"100%",width:`${eval0.avgScore*20}%`,background:overallColor(eval0.avgScore),borderRadius:3}}/>
            </div>
            <p style={{fontSize:10,color:"#94a3b8",margin:"8px 0 0"}}>Média de {completed.length} testes realizados</p>
          </div>

          {/* Individual Results */}
          {completed.map(([testId, result]) => {
            const test = TESTS.find(t => t.id === testId);
            return (
              <div key={testId} style={{...S.resultCard, borderLeftColor: result.color}}>
                <div style={{display:"flex",justifyContent:"space-between",alignItems:"center"}}>
                  <div>
                    <p style={{fontSize:14,fontWeight:700,color:"#e2e8f0",margin:0}}>{test.icon} {test.name}</p>
                    <p style={{fontSize:11,color:"#94a3b8",margin:"2px 0 0"}}>{test.desc.split("(")[1]?.replace(")","") || ""}</p>
                  </div>
                  <div style={{textAlign:"right"}}>
                    <span style={{fontSize:22,fontWeight:900,color:result.color}}>{result.value}</span>
                    <span style={{display:"block",fontSize:10,color:"#94a3b8"}}>{test.unit}</span>
                  </div>
                </div>
                <div style={{display:"flex",alignItems:"center",gap:8,marginTop:8}}>
                  <span style={{fontSize:16}}>{result.emoji}</span>
                  <span style={{fontSize:13,fontWeight:700,color:result.color}}>{result.class}</span>
                  {result.percentile && <span style={{fontSize:10,color:"#64748b"}}>P{result.percentile}</span>}
                  {result.tScale && <span style={{fontSize:10,color:"#64748b"}}>T-Scale: {result.tScale}</span>}
                  <div style={{flex:1}}/>
                  <span style={{fontSize:12,fontWeight:700,color:result.color}}>{result.score}/5</span>
                </div>
              </div>
            );
          })}

          {/* Disclaimer */}
          <div style={S.disclaimerBox}>
            <p style={{fontSize:11,color:"#94a3b8",margin:0,lineHeight:1.6}}>
              ⚕️ Esta avaliação é baseada em tabelas normativas publicadas em periódicos científicos indexados. Os resultados são indicativos e não substituem uma avaliação profissional completa. Consulte um profissional de Educação Física ou saúde para prescrição de exercícios.
            </p>
          </div>

          <button style={{...S.primaryBtn,marginTop:12}} onClick={() => {setResults({abdominal:null,pushup:null,plank:null,burpee:null});setScreen("tests")}}>
            🔄 Nova Avaliação
          </button>
        </div>
      </div>
    );
  }

  // ─── HISTORY ───────────────────────────────────────────────
  if (screen === "history") {
    return (
      <div style={S.app}>
        <div style={S.header}>
          <button style={S.backBtn} onClick={() => setScreen("home")}>← Voltar</button>
          <h2 style={S.headerTitle}>📊 Histórico</h2>
          <div style={{width:50}}/>
        </div>
        <div style={{padding:"0 20px"}}>
          {evaluations.map((ev, i) => {
            const ct = Object.values(ev.results).filter(r => r !== null).length;
            return (
              <div key={i} style={S.historyCard}>
                <div style={{display:"flex",justifyContent:"space-between",alignItems:"center"}}>
                  <div>
                    <p style={{fontSize:13,fontWeight:700,color:"#e2e8f0",margin:0}}>{ev.profile.name || "Avaliado"}</p>
                    <p style={{fontSize:11,color:"#94a3b8",margin:"2px 0 0"}}>{ev.date} • {ct} testes</p>
                  </div>
                  <div style={{textAlign:"right"}}>
                    <span style={{fontSize:24,fontWeight:900,color:overallColor(ev.avgScore)}}>{ev.avgScore.toFixed(1)}</span>
                    <span style={{display:"block",fontSize:10,color:overallColor(ev.avgScore),fontWeight:600}}>{ev.overallClass}</span>
                  </div>
                </div>
                <div style={{display:"flex",gap:6,flexWrap:"wrap",marginTop:8}}>
                  {Object.entries(ev.results).filter(([,v])=>v).map(([id,r]) => (
                    <span key={id} style={{fontSize:10,padding:"3px 8px",borderRadius:6,border:`1px solid ${r.color}`,color:r.color,fontWeight:600}}>
                      {r.emoji} {TESTS.find(t=>t.id===id)?.name.split(" ")[0]} {r.value}
                    </span>
                  ))}
                </div>
              </div>
            );
          })}
        </div>
      </div>
    );
  }

  return null;
}

// ═══════════════════════════════════════════════════════════════
// STYLES
// ═══════════════════════════════════════════════════════════════
const S = {
  app: { minHeight:"100vh",background:"linear-gradient(175deg,#0a0e1a 0%,#0f1629 50%,#0a1220 100%)",color:"#e2e8f0",fontFamily:"'Segoe UI',-apple-system,BlinkMacSystemFont,sans-serif",maxWidth:500,margin:"0 auto",paddingBottom:40 },
  homeWrap: { padding:"24px 20px" },
  brand: { textAlign:"center",marginBottom:28 },
  brandIcon: { display:"inline-block",marginBottom:8,filter:"drop-shadow(0 0 24px rgba(99,102,241,0.25))" },
  brandTitle: { fontSize:36,fontWeight:900,margin:0,background:"linear-gradient(135deg,#6366f1,#10b981)",WebkitBackgroundClip:"text",WebkitTextFillColor:"transparent",letterSpacing:-1 },
  brandSub: { fontSize:14,color:"#94a3b8",margin:"4px 0 0",fontWeight:500 },
  brandRef: { fontSize:11,color:"#64748b",margin:"2px 0 0" },

  profileCard: { background:"rgba(15,20,40,0.7)",borderRadius:16,padding:"20px 18px",border:"1px solid rgba(99,102,241,0.12)",marginBottom:20 },
  sectionTitle: { fontSize:15,fontWeight:700,color:"#c7d2fe",margin:"0 0 14px" },
  formGrid: { display:"flex",flexDirection:"column",gap:12 },
  field: {},
  fieldRow: { display:"flex",gap:12 },
  label: { fontSize:11,fontWeight:600,color:"#94a3b8",display:"block",marginBottom:4,textTransform:"uppercase",letterSpacing:0.5 },
  input: { width:"100%",padding:"10px 12px",fontSize:14,background:"rgba(30,35,60,0.8)",border:"1px solid rgba(100,116,139,0.2)",borderRadius:10,color:"#e2e8f0",outline:"none",boxSizing:"border-box" },
  segmented: { display:"flex",gap:0,borderRadius:10,overflow:"hidden",border:"1px solid rgba(100,116,139,0.2)" },
  segBtn: { flex:1,padding:"10px 0",fontSize:13,fontWeight:600,background:"rgba(30,35,60,0.8)",border:"none",color:"#94a3b8",cursor:"pointer",textAlign:"center" },
  segBtnActive: { background:"rgba(99,102,241,0.25)",color:"#a5b4fc" },
  checkRow: { display:"flex",alignItems:"center" },
  checkLabel: { display:"flex",alignItems:"center",gap:8,fontSize:13,color:"#c7d2fe" },
  checkbox: { width:20,height:20,borderRadius:6,border:"2px solid rgba(100,116,139,0.3)",display:"flex",alignItems:"center",justifyContent:"center",fontSize:12,fontWeight:700,color:"#6366f1",flexShrink:0 },
  checkboxActive: { background:"rgba(99,102,241,0.2)",borderColor:"#6366f1" },

  primaryBtn: { display:"flex",alignItems:"center",justifyContent:"center",gap:10,width:"100%",padding:"15px 24px",fontSize:15,fontWeight:700,color:"#fff",background:"linear-gradient(135deg,#6366f1,#4f46e5)",border:"none",borderRadius:14,cursor:"pointer",boxShadow:"0 4px 20px rgba(99,102,241,0.25)",marginBottom:10 },
  ghostBtn: { display:"flex",alignItems:"center",justifyContent:"center",gap:8,width:"100%",padding:"12px 20px",fontSize:14,fontWeight:600,color:"#a5b4fc",background:"transparent",border:"1px solid rgba(99,102,241,0.2)",borderRadius:12,cursor:"pointer" },

  refs: { marginTop:24,padding:"16px",background:"rgba(10,15,30,0.5)",borderRadius:12,border:"1px solid rgba(100,116,139,0.08)" },
  refTitle: { fontSize:12,fontWeight:700,color:"#94a3b8",margin:"0 0 6px" },
  refText: { fontSize:11,color:"#64748b",margin:"3px 0",lineHeight:1.4 },

  header: { display:"flex",alignItems:"center",justifyContent:"space-between",padding:"14px 18px",borderBottom:"1px solid rgba(100,116,139,0.12)",background:"rgba(10,15,30,0.9)",backdropFilter:"blur(10px)",position:"sticky",top:0,zIndex:10 },
  headerTitle: { fontSize:15,fontWeight:700,margin:0,color:"#c7d2fe" },
  backBtn: { background:"none",border:"none",color:"#818cf8",fontSize:13,fontWeight:600,cursor:"pointer",padding:"4px 8px" },
  badge: { fontSize:11,fontWeight:700,color:"#a5b4fc",background:"rgba(99,102,241,0.15)",padding:"4px 10px",borderRadius:20 },

  testListWrap: { padding:"12px 18px" },
  profileMini: { display:"flex",justifyContent:"space-between",alignItems:"center",padding:"10px 14px",background:"rgba(20,25,45,0.6)",borderRadius:10,marginBottom:14,fontSize:13,fontWeight:600,color:"#e2e8f0" },
  profileTag: { fontSize:11,color:"#94a3b8",fontWeight:500 },

  testCard: { padding:"16px",background:"rgba(15,20,38,0.7)",borderRadius:14,border:"1px solid rgba(100,116,139,0.1)",marginBottom:10 },
  testCardTop: { display:"flex",alignItems:"center",gap:14 },
  testName: { fontSize:15,fontWeight:700,color:"#e2e8f0",margin:0 },
  testDesc: { fontSize:11,color:"#94a3b8",margin:"3px 0 0" },
  startBtn: { padding:"8px 18px",fontSize:12,fontWeight:700,background:"rgba(99,102,241,0.15)",border:"1px solid rgba(99,102,241,0.3)",borderRadius:10,color:"#a5b4fc",cursor:"pointer",flexShrink:0 },
  redoBtn: { fontSize:12,background:"none",border:"1px solid rgba(100,116,139,0.2)",borderRadius:6,color:"#94a3b8",cursor:"pointer",padding:"2px 8px" },

  protocolCard: { background:"rgba(99,102,241,0.06)",borderRadius:12,padding:"14px 16px",border:"1px solid rgba(99,102,241,0.1)",marginBottom:16 },
  timerCard: { textAlign:"center",padding:"24px 16px",background:"rgba(15,20,38,0.8)",borderRadius:16,border:"1px solid rgba(100,116,139,0.1)",marginBottom:16 },
  timerDisplay: { fontSize:48,fontWeight:900,fontFamily:"'Courier New',monospace",color:"#e2e8f0",letterSpacing:2,marginBottom:14 },
  timerBtn: { padding:"10px 20px",fontSize:13,fontWeight:700,background:"rgba(99,102,241,0.15)",border:"1px solid rgba(99,102,241,0.3)",borderRadius:10,color:"#a5b4fc",cursor:"pointer" },
  timerBtnSmall: { padding:"6px 14px",fontSize:11,fontWeight:600,background:"none",border:"1px solid rgba(100,116,139,0.2)",borderRadius:8,color:"#94a3b8",cursor:"pointer" },

  inputCard: { background:"rgba(15,20,38,0.6)",borderRadius:14,padding:"16px",border:"1px solid rgba(100,116,139,0.1)",marginBottom:16 },
  tableCard: { background:"rgba(10,15,30,0.6)",borderRadius:12,padding:"14px 16px",border:"1px solid rgba(100,116,139,0.08)" },

  scoreCard: { textAlign:"center",padding:"24px 20px 18px",borderRadius:20,background:"linear-gradient(170deg,rgba(15,20,40,0.9),rgba(10,15,30,0.9))",border:"2px solid",marginBottom:20,boxShadow:"0 8px 32px rgba(0,0,0,0.3)" },
  reportHeader: { display:"flex",gap:14,alignItems:"center",marginBottom:16,padding:"16px",background:"rgba(15,20,38,0.7)",borderRadius:14,border:"1px solid rgba(99,102,241,0.1)" },
  reportLogo: { flexShrink:0 },
  reportProfile: { background:"rgba(15,20,38,0.6)",borderRadius:12,padding:"14px 16px",marginBottom:16,border:"1px solid rgba(100,116,139,0.08)" },
  rpRow: { display:"flex",justifyContent:"space-between",padding:"4px 0",borderBottom:"1px solid rgba(100,116,139,0.06)" },
  rpLabel: { fontSize:12,color:"#94a3b8" },
  rpVal: { fontSize:12,color:"#e2e8f0",fontWeight:600 },

  resultCard: { padding:"14px 16px",background:"rgba(15,20,38,0.6)",borderRadius:14,borderLeft:"4px solid",marginBottom:10 },

  disclaimerBox: { padding:"14px 16px",background:"rgba(234,179,8,0.05)",borderRadius:12,border:"1px solid rgba(234,179,8,0.1)",marginTop:16 },

  historyCard: { padding:"14px 16px",background:"rgba(15,20,38,0.6)",borderRadius:14,border:"1px solid rgba(100,116,139,0.08)",marginBottom:10 },
};

if (typeof document !== "undefined" && !document.getElementById("fiteval-styles")) {
  const style = document.createElement("style");
  style.id = "fiteval-styles";
  style.textContent = `
    * { box-sizing: border-box; -webkit-tap-highlight-color: transparent; }
    body { margin: 0; background: #0a0e1a; }
    button:hover { filter: brightness(1.08); }
    button:active { transform: scale(0.97); }
    input:focus { border-color: #6366f1 !important; box-shadow: 0 0 0 2px rgba(99,102,241,0.15); }
    input[type=number]::-webkit-inner-spin-button { opacity: 1; }
  `;
  document.head.appendChild(style);
}
