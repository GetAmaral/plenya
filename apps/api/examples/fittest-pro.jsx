import { useState, useRef, useEffect } from "react";

// ═══════════════════════════════════════════════════════════
// NORMATIVE DATA FROM UPLOADED RESEARCH PAPERS
// ═══════════════════════════════════════════════════════════

// ABDOMINAL - reps/min (Pollock & Wilmore, 1993) — Images 1 & 3
const AB = {
  male: {
    "15-19": [[48,999],[42,47],[38,41],[33,37],[0,32]],
    "20-29": [[43,999],[37,42],[33,36],[29,32],[0,28]],
    "30-39": [[36,999],[31,35],[27,30],[22,26],[0,21]],
    "40-49": [[31,999],[26,30],[22,25],[17,21],[0,16]],
    "50-59": [[26,999],[22,25],[18,21],[13,17],[0,12]],
    "60-69": [[23,999],[17,22],[12,16],[7,11],[0,6]],
  },
  female: {
    "15-19": [[42,999],[36,41],[32,35],[27,31],[0,26]],
    "20-29": [[36,999],[31,35],[25,30],[21,24],[0,20]],
    "30-39": [[29,999],[24,28],[20,23],[15,19],[0,14]],
    "40-49": [[25,999],[20,24],[15,19],[7,14],[0,6]],
    "50-59": [[19,999],[12,18],[5,11],[3,4],[0,2]],
    "60-69": [[16,999],[12,15],[4,11],[2,3],[0,1]],
  }
};

// PUSH-UP - reps/min (Pollock & Wilmore, 1993) — Image 2
const PU = {
  male: {
    "15-19": [[39,999],[29,38],[23,28],[18,22],[0,17]],
    "20-29": [[36,999],[29,35],[22,28],[17,21],[0,16]],
    "30-39": [[30,999],[22,29],[17,21],[12,16],[0,11]],
    "40-49": [[22,999],[17,21],[13,16],[10,12],[0,9]],
    "50-59": [[21,999],[13,20],[10,12],[7,9],[0,6]],
    "60-69": [[18,999],[11,17],[8,10],[5,7],[0,4]],
  },
  female: {
    "15-19": [[33,999],[25,32],[18,24],[12,17],[0,11]],
    "20-29": [[30,999],[21,29],[15,20],[10,14],[0,9]],
    "30-39": [[27,999],[20,26],[13,19],[8,12],[0,7]],
    "40-49": [[24,999],[15,23],[11,14],[5,10],[0,4]],
    "50-59": [[21,999],[11,20],[7,10],[2,6],[0,1]],
    "60-69": [[17,999],[12,16],[5,11],[2,4],[0,1]],
  }
};

// PLANK - seconds (Strand et al., JHK 40, 2014) — percentiles
const PK = {
  male: [[201,9999],[157,200],[97,156],[62,96],[0,61]],
  female: [[142,9999],[108,141],[63,107],[35,62],[0,34]],
};

// 3-MIN BURPEE (Podstawski et al., JHK 69, 2019) — cycles/3min
const BU = {
  male: [[76,85],[67,75],[47,66],[38,46],[0,37]],
  female: [[72,83],[61,71],[37,60],[26,36],[0,25]],
};

// FRT - Flexion Rotation Trunk (Brotons-Gil et al., JSCR 2013) — reps/90s
// Using T4 stabilized data: Males ~100±26, Females ~72±18
const FRT_STATS = { male: {m:100,s:26}, female: {m:72,s:18} };

const LEVELS = ["excelente","bom","medio","abaixo","fraco"];
const LV = {
  excelente: {l:"Excelente",c:"#22c55e",e:"🏆",bg:"rgba(34,197,94,0.08)",bd:"rgba(34,197,94,0.25)"},
  bom:       {l:"Bom",c:"#3b82f6",e:"👍",bg:"rgba(59,130,246,0.08)",bd:"rgba(59,130,246,0.25)"},
  medio:     {l:"Médio",c:"#eab308",e:"➡️",bg:"rgba(234,179,8,0.08)",bd:"rgba(234,179,8,0.25)"},
  abaixo:    {l:"Abaixo da Média",c:"#f97316",e:"⚠️",bg:"rgba(249,115,22,0.08)",bd:"rgba(249,115,22,0.25)"},
  fraco:     {l:"Fraco",c:"#ef4444",e:"🔴",bg:"rgba(239,68,68,0.08)",bd:"rgba(239,68,68,0.25)"},
};
const SCORES = {excelente:5,bom:4,medio:3,abaixo:2,fraco:1};

function ageGroup(age) {
  if(age>=15&&age<=19) return "15-19";
  if(age>=20&&age<=29) return "20-29";
  if(age>=30&&age<=39) return "30-39";
  if(age>=40&&age<=49) return "40-49";
  if(age>=50&&age<=59) return "50-59";
  if(age>=60&&age<=69) return "60-69";
  return null;
}

function classify(val, ranges) {
  for (let i=0;i<ranges.length;i++) {
    if(val>=ranges[i][0] && val<=ranges[i][1]) return LEVELS[i];
  }
  return "fraco";
}

function classifyFRT(val, gender) {
  const {m,s} = FRT_STATS[gender];
  const z = (val-m)/s;
  if(z>=1.5) return "excelente";
  if(z>=0.5) return "bom";
  if(z>=-0.5) return "medio";
  if(z>=-1.5) return "abaixo";
  return "fraco";
}

const TESTS = [
  {id:"abdominal",name:"Resistência Abdominal",icon:"🔥",unit:"rep/min",time:60,ages:"15-69",
   desc:"Máximo de abdominais em 1 minuto",ref:"Pollock & Wilmore, 1993"},
  {id:"pushup",name:"Flexão de Braço",icon:"💪",unit:"rep/min",time:60,ages:"15-69",
   desc:"Máximo de flexões de solo em 1 minuto",ref:"Pollock & Wilmore, 1993"},
  {id:"plank",name:"Prancha Isométrica",icon:"🧘",unit:"segundos",time:0,ages:"18-29",
   desc:"Manter prancha ventral até a falha",ref:"Strand et al., JHK, 2014"},
  {id:"burpee",name:"Burpee 3 Minutos",icon:"🏃",unit:"ciclos/3min",time:180,ages:"18-25",
   desc:"Máximo de ciclos de burpee em 3 minutos — Resistência cardiorrespiratória",ref:"Podstawski et al., JHK, 2019"},
  {id:"frt",name:"Flexão-Rotação do Tronco",icon:"🔄",unit:"rep/90s",time:90,ages:"18-30",
   desc:"Máximo de flexões-rotações alternadas em 90 segundos",ref:"Brotons-Gil et al., JSCR, 2013"},
];

function evaluate(testId, val, gender, age_) {
  const v = Number(val);
  if(isNaN(v)||v<0) return null;
  const g = gender;
  const ag = ageGroup(age_);
  let cl = null, ref = "";

  if(testId==="abdominal") {
    if(!ag||!AB[g]?.[ag]) return null;
    cl = classify(v, AB[g][ag]);
    ref = "Pollock & Wilmore, 1993";
  } else if(testId==="pushup") {
    if(!ag||!PU[g]?.[ag]) return null;
    cl = classify(v, PU[g][ag]);
    ref = "Pollock & Wilmore, 1993";
  } else if(testId==="plank") {
    cl = classify(v, PK[g]);
    ref = "Strand et al., JHK, 2014";
  } else if(testId==="burpee") {
    cl = classify(v, BU[g]);
    ref = "Podstawski et al., JHK, 2019";
  } else if(testId==="frt") {
    cl = classifyFRT(v, g);
    ref = "Brotons-Gil et al., JSCR, 2013";
  }
  return {value:v, cl, ref};
}

// ═══════════════════════════════════════════════════════════
// APP
// ═══════════════════════════════════════════════════════════

export default function App() {
  const [scr, setScr] = useState("home");
  const [sex, setSex] = useState("male");
  const [age, setAge] = useState(25);
  const [res, setRes] = useState({});
  const [cur, setCur] = useState(null);
  const [inp, setInp] = useState("");
  const [tRun, setTRun] = useState(false);
  const [tSec, setTSec] = useState(0);
  const [tMode, setTMode] = useState("cd");
  const tRef = useRef(null);

  useEffect(() => {
    if(tRun) {
      tRef.current = setInterval(() => {
        setTSec(p => {
          if(tMode==="cd") { if(p<=1){setTRun(false);return 0;} return p-1; }
          return p+1;
        });
      },1000);
    }
    return () => clearInterval(tRef.current);
  },[tRun,tMode]);

  const startT = (s,m) => {setTMode(m);setTSec(m==="cd"?s:0);setTRun(true);};
  const stopT = () => {setTRun(false);clearInterval(tRef.current);};
  const fmtT = s => `${Math.floor(s/60)}:${(s%60).toString().padStart(2,"0")}`;

  const save = () => {
    let v = inp || (cur==="plank"&&!tRun&&tSec>0? String(tSec):"");
    if(!v) return;
    const ev = evaluate(cur, v, sex, age);
    if(ev) { setRes(p=>({...p,[cur]:ev})); setInp(""); stopT(); setTSec(0); setScr("result"); }
  };

  const overall = () => {
    const k=Object.keys(res);
    if(!k.length) return null;
    return Math.round(k.reduce((s,t)=>s+(SCORES[res[t].cl]||0),0)/(k.length*5)*100);
  };
  const ovLabel = s => s>=85?{l:"Excelente",c:"#22c55e"}:s>=70?{l:"Bom",c:"#3b82f6"}:s>=50?{l:"Médio",c:"#eab308"}:s>=30?{l:"Regular",c:"#f97316"}:{l:"Atenção",c:"#ef4444"};

  // HOME
  if(scr==="home") return (
    <div style={Z.app}><div style={Z.ct}>
      <div style={{textAlign:"center",padding:"24px 0 16px"}}>
        <div style={{fontSize:44}}>🏋️</div>
        <h1 style={Z.title}>FitTest Pro</h1>
        <p style={{fontSize:12,color:"#94a3b8",margin:"4px 0 0"}}>Avaliação Física Baseada em Evidências</p>
      </div>

      <div style={Z.card}>
        <p style={{fontSize:14,fontWeight:700,color:"#c7d2fe",margin:"0 0 12px"}}>Dados do Avaliado</p>
        <div style={{marginBottom:12}}>
          <span style={Z.lbl}>Sexo</span>
          <div style={{display:"flex",gap:8}}>
            {[["male","♂ Masculino"],["female","♀ Feminino"]].map(([v,l])=>(
              <button key={v} style={{...Z.tog,...(sex===v?Z.togA:{})}} onClick={()=>setSex(v)}>{l}</button>
            ))}
          </div>
        </div>
        <div style={{marginBottom:12}}>
          <span style={Z.lbl}>Idade: <strong style={{color:"#e2e8f0"}}>{age}</strong> anos</span>
          <input type="range" min={15} max={69} value={age} onChange={e=>setAge(+e.target.value)} style={{width:"100%",accentColor:"#6366f1"}}/>
        </div>
      </div>

      <div style={Z.card}>
        <p style={{fontSize:14,fontWeight:700,color:"#c7d2fe",margin:"0 0 10px"}}>Testes</p>
        {TESTS.map(t => {
          const d = res[t.id];
          const cf = d? LV[d.cl]:null;
          return (
            <div key={t.id} style={{...Z.ti,borderColor:d?cf?.bd:"rgba(100,116,139,0.12)"}} onClick={()=>{setCur(t.id);setScr("input")}}>
              <div style={{display:"flex",alignItems:"center",gap:10}}>
                <span style={{fontSize:26}}>{t.icon}</span>
                <div>
                  <p style={{fontSize:13,fontWeight:700,color:"#e2e8f0",margin:0}}>{t.name}</p>
                  <p style={{fontSize:10,color:"#94a3b8",margin:"2px 0 0"}}>{t.unit} • {t.ages} anos</p>
                </div>
              </div>
              {d ? <div style={{textAlign:"right"}}><span style={{fontSize:16,fontWeight:800,color:cf.c}}>{d.value}</span><span style={{display:"block",fontSize:9,color:cf.c,fontWeight:600}}>{cf.e} {cf.l}</span></div>
                 : <span style={{fontSize:10,color:"#64748b",padding:"4px 10px",border:"1px solid rgba(100,116,139,0.2)",borderRadius:8}}>Avaliar</span>}
            </div>
          );
        })}
      </div>

      {Object.keys(res).length>0 && (
        <button style={Z.pri} onClick={()=>setScr("report")}>📊 Relatório ({Object.keys(res).length}/{TESTS.length})</button>
      )}

      <div style={{marginTop:16,padding:"12px 14px",background:"rgba(15,15,30,0.3)",borderRadius:10,border:"1px solid rgba(100,116,139,0.06)"}}>
        <p style={{fontSize:10,fontWeight:700,color:"#64748b",margin:"0 0 4px"}}>📚 Referências</p>
        {["Pollock & Wilmore (1993)","Strand et al., JHK 40, 2014","Podstawski et al., JHK 69, 2019","Adams et al., IJES 15(4), 2022","Brotons-Gil et al., JSCR 27(6), 2013"].map((r,i)=>(
          <p key={i} style={{fontSize:9,color:"#475569",margin:"2px 0"}}>{r}</p>
        ))}
      </div>
    </div></div>
  );

  // INPUT
  if(scr==="input"&&cur) {
    const t = TESTS.find(x=>x.id===cur);
    const isPlank = cur==="plank";
    return (
      <div style={Z.app}>
        <div style={Z.hd}>
          <button style={Z.bk} onClick={()=>{stopT();setTSec(0);setScr("home")}}>← Voltar</button>
          <h2 style={Z.ht}>{t.icon} {t.name}</h2>
          <div style={{width:50}}/>
        </div>
        <div style={Z.ct}>
          <div style={Z.card}>
            <p style={{fontSize:13,color:"#c7d2fe",margin:"0 0 6px",lineHeight:1.5}}>{t.desc}</p>
            <p style={{fontSize:10,color:"#64748b",margin:0}}>Ref: {t.ref}</p>
          </div>

          {/* Timer */}
          <div style={{textAlign:"center",padding:"20px",marginBottom:14,background:"rgba(15,15,35,0.5)",borderRadius:16,border:"1px solid rgba(99,102,241,0.12)"}}>
            <div style={{fontSize:52,fontWeight:900,color:"#e2e8f0",fontVariantNumeric:"tabular-nums",lineHeight:1,marginBottom:10}}>{fmtT(tSec)}</div>
            {tRun && <div style={{height:5,background:"rgba(100,116,139,0.12)",borderRadius:3,marginBottom:12,overflow:"hidden"}}>
              <div style={{height:"100%",background:"linear-gradient(90deg,#6366f1,#22d3ee)",borderRadius:3,transition:"width 0.3s",
                width: tMode==="cd"?`${((t.time-tSec)/t.time)*100}%`:"100%",
                ...(tMode==="sw"?{animation:"pulse 1.5s infinite"}:{})
              }}/>
            </div>}
            <div style={{display:"flex",gap:8,justifyContent:"center"}}>
              {!tRun ? <>
                {t.time>0 && <button style={Z.tb} onClick={()=>startT(t.time,"cd")}>▶️ {t.time}s</button>}
                {isPlank && <button style={Z.tb} onClick={()=>startT(0,"sw")}>⏱ Cronômetro</button>}
                {!isPlank && t.time===0 && <button style={Z.tb} onClick={()=>startT(0,"sw")}>⏱ Iniciar</button>}
              </> : <button style={{...Z.tb,background:"rgba(239,68,68,0.15)",borderColor:"rgba(239,68,68,0.3)"}} onClick={stopT}>⏹ Parar{isPlank?` — ${tSec}s`:""}</button>}
            </div>
          </div>

          <div style={Z.card}>
            <span style={{...Z.lbl,marginBottom:8,display:"block"}}>Resultado ({t.unit})</span>
            <div style={{display:"flex",gap:8}}>
              <input type="number" inputMode="numeric" value={inp||(isPlank&&!tRun&&tSec>0?tSec:"")} onChange={e=>setInp(e.target.value)}
                placeholder={cur==="plank"?"seg":cur==="burpee"?"ciclos":"reps"} style={Z.inp}/>
              <button style={{...Z.pri,flex:"none",padding:"12px 20px"}} onClick={()=>{if(!inp&&isPlank&&tSec>0)setInp(String(tSec));save();}}>Salvar</button>
            </div>
            {inp && (()=>{
              const ev=evaluate(cur,inp,sex,age);
              if(!ev) return <p style={{fontSize:11,color:"#f87171",marginTop:8}}>Faixa etária fora das normas.</p>;
              const cf=LV[ev.cl];
              return <div style={{marginTop:10,padding:"8px 14px",background:cf.bg,borderRadius:10,border:`1px solid ${cf.bd}`}}>
                <span style={{fontSize:13,fontWeight:700,color:cf.c}}>{cf.e} {cf.l}</span></div>;
            })()}
          </div>

          {/* Normative table preview */}
          {(cur==="abdominal"||cur==="pushup") && (()=>{
            const norms = cur==="abdominal"?AB:PU;
            const ag = ageGroup(age);
            if(!ag||!norms[sex]?.[ag]) return null;
            const ranges = norms[sex][ag];
            return <div style={Z.card}>
              <p style={{fontSize:11,fontWeight:700,color:"#94a3b8",margin:"0 0 8px"}}>📋 Norma — {sex==="male"?"Homens":"Mulheres"}, {ag} anos</p>
              {LEVELS.map((lv,i)=>{
                const cf=LV[lv];const [mn,mx]=ranges[i];
                return <div key={lv} style={{display:"flex",justifyContent:"space-between",padding:"5px 0",borderBottom:"1px solid rgba(100,116,139,0.06)"}}>
                  <span style={{fontSize:11,color:cf.c,fontWeight:600}}>{cf.l}</span>
                  <span style={{fontSize:11,color:"#94a3b8"}}>{mx>=999?`≥ ${mn}`:`${mn} – ${mx}`}</span>
                </div>;
              })}
            </div>;
          })()}
        </div>
      </div>
    );
  }

  // RESULT
  if(scr==="result"&&cur) {
    const t=TESTS.find(x=>x.id===cur);
    const r=res[cur];
    const cf=LV[r.cl];
    return (
      <div style={Z.app}>
        <div style={Z.hd}>
          <button style={Z.bk} onClick={()=>setScr("home")}>← Testes</button>
          <h2 style={Z.ht}>{t.icon} Resultado</h2>
          <div style={{width:50}}/>
        </div>
        <div style={Z.ct}>
          <div style={{textAlign:"center",padding:"32px 20px",marginBottom:16,borderRadius:20,background:"rgba(15,15,35,0.7)",border:`1px solid ${cf.bd}`}}>
            <span style={{fontSize:52,fontWeight:900,color:cf.c}}>{r.value}</span>
            <span style={{display:"block",fontSize:13,color:"#94a3b8",marginTop:2}}>{t.unit}</span>
            <div style={{marginTop:14,padding:"8px 24px",display:"inline-block",background:cf.bg,borderRadius:12,border:`1px solid ${cf.bd}`}}>
              <span style={{fontSize:18,fontWeight:700,color:cf.c}}>{cf.e} {cf.l}</span>
            </div>
            <p style={{fontSize:11,color:"#64748b",marginTop:10}}>{sex==="male"?"Masculino":"Feminino"}, {age} anos</p>
            <p style={{fontSize:9,color:"#475569",marginTop:2}}>Ref: {r.ref}</p>
          </div>
          <div style={{display:"flex",gap:10}}>
            <button style={{...Z.sec,flex:1}} onClick={()=>{setCur(null);setScr("home")}}>← Outros Testes</button>
            {Object.keys(res).length>=2 && <button style={{...Z.pri,flex:1}} onClick={()=>setScr("report")}>📊 Relatório</button>}
          </div>
        </div>
      </div>
    );
  }

  // REPORT
  if(scr==="report") {
    const ov=overall();
    const ovl=ov!==null?ovLabel(ov):null;
    return (
      <div style={Z.app}>
        <div style={Z.hd}>
          <button style={Z.bk} onClick={()=>setScr("home")}>← Voltar</button>
          <h2 style={Z.ht}>📊 Relatório</h2>
          <div style={{width:50}}/>
        </div>
        <div style={Z.ct}>
          <div style={{textAlign:"center",padding:"16px 0 12px"}}>
            <h2 style={{fontSize:18,fontWeight:800,color:"#e2e8f0",margin:0}}>Avaliação Física Completa</h2>
            <p style={{fontSize:12,color:"#94a3b8",margin:"4px 0 0"}}>{sex==="male"?"♂ Masculino":"♀ Feminino"} • {age} anos • {new Date().toLocaleDateString("pt-BR")}</p>
          </div>

          {ov!==null && <div style={{textAlign:"center",padding:"24px 20px",marginBottom:16,borderRadius:18,background:"rgba(15,15,35,0.6)",border:"1px solid rgba(100,116,139,0.1)"}}>
            <span style={{fontSize:52,fontWeight:900,color:ovl.c}}>{ov}</span>
            <span style={{display:"block",fontSize:10,color:"#94a3b8",marginTop:2}}>SCORE GERAL</span>
            <div style={{marginTop:6,display:"inline-block",padding:"4px 16px",borderRadius:8,background:`${ovl.c}18`}}>
              <span style={{fontSize:14,fontWeight:700,color:ovl.c}}>{ovl.l}</span>
            </div>
            <div style={{width:"100%",height:5,background:"rgba(100,116,139,0.12)",borderRadius:3,marginTop:12,overflow:"hidden"}}>
              <div style={{height:"100%",width:`${ov}%`,background:ovl.c,borderRadius:3,transition:"width 0.8s"}}/>
            </div>
          </div>}

          {Object.entries(res).map(([tid,r])=>{
            const t=TESTS.find(x=>x.id===tid);
            const cf=LV[r.cl];
            return <div key={tid} style={{padding:"14px 16px",background:"rgba(20,20,42,0.4)",borderRadius:14,border:`1px solid ${cf.bd}`,marginBottom:10}}>
              <div style={{display:"flex",justifyContent:"space-between",alignItems:"center"}}>
                <div style={{display:"flex",alignItems:"center",gap:10}}>
                  <span style={{fontSize:22}}>{t.icon}</span>
                  <div><p style={{fontSize:13,fontWeight:700,color:"#e2e8f0",margin:0}}>{t.name}</p><p style={{fontSize:10,color:"#94a3b8",margin:"2px 0 0"}}>{t.unit}</p></div>
                </div>
                <div style={{textAlign:"right"}}>
                  <span style={{fontSize:20,fontWeight:800,color:cf.c}}>{r.value}</span>
                  <span style={{display:"block",fontSize:10,fontWeight:600,color:cf.c}}>{cf.e} {cf.l}</span>
                </div>
              </div>
              <div style={{height:3,background:"rgba(100,116,139,0.08)",borderRadius:2,marginTop:8,overflow:"hidden"}}>
                <div style={{height:"100%",width:`${(SCORES[r.cl]/5)*100}%`,background:cf.c,borderRadius:2}}/>
              </div>
              <p style={{fontSize:9,color:"#475569",margin:"5px 0 0"}}>Ref: {r.ref}</p>
            </div>;
          })}

          {Object.keys(res).length<TESTS.length && <div style={{textAlign:"center",padding:12}}>
            <p style={{fontSize:12,color:"#94a3b8",margin:"0 0 8px"}}>{TESTS.length-Object.keys(res).length} teste(s) pendente(s)</p>
            <button style={Z.sec} onClick={()=>setScr("home")}>Continuar Avaliação</button>
          </div>}

          <p style={{fontSize:10,color:"#64748b",textAlign:"center",marginTop:16,lineHeight:1.5}}>⚕️ Resultados baseados em normas científicas. Interpretação por profissional qualificado.</p>
        </div>
      </div>
    );
  }

  return null;
}

// STYLES
const Z = {
  app:{minHeight:"100vh",background:"linear-gradient(170deg,#0c0c1d,#111128 50%,#0a1018)",color:"#e2e8f0",fontFamily:"'Segoe UI',-apple-system,sans-serif",maxWidth:500,margin:"0 auto",paddingBottom:36},
  ct:{padding:"14px 18px"},
  title:{fontSize:28,fontWeight:900,margin:"6px 0 0",background:"linear-gradient(135deg,#818cf8,#22d3ee)",WebkitBackgroundClip:"text",WebkitTextFillColor:"transparent"},
  card:{background:"rgba(20,20,42,0.5)",borderRadius:14,padding:"16px",border:"1px solid rgba(100,116,139,0.1)",marginBottom:14},
  lbl:{fontSize:11,fontWeight:600,color:"#94a3b8",marginBottom:4,display:"block"},
  tog:{flex:1,padding:"10px",fontSize:12,fontWeight:600,background:"rgba(30,30,55,0.5)",border:"1px solid rgba(100,116,139,0.15)",borderRadius:10,color:"#94a3b8",cursor:"pointer",textAlign:"center"},
  togA:{background:"rgba(99,102,241,0.12)",borderColor:"#6366f1",color:"#c7d2fe"},
  ti:{display:"flex",justifyContent:"space-between",alignItems:"center",padding:"12px 14px",background:"rgba(15,15,35,0.3)",borderRadius:12,border:"1px solid rgba(100,116,139,0.1)",marginBottom:8,cursor:"pointer"},
  pri:{display:"flex",alignItems:"center",justifyContent:"center",gap:8,padding:"14px 22px",fontSize:14,fontWeight:700,color:"#fff",background:"linear-gradient(135deg,#6366f1,#4f46e5)",border:"none",borderRadius:14,cursor:"pointer",width:"100%"},
  sec:{display:"flex",alignItems:"center",justifyContent:"center",gap:6,padding:"11px 18px",fontSize:12,fontWeight:600,color:"#c7d2fe",background:"rgba(99,102,241,0.08)",border:"1px solid rgba(99,102,241,0.18)",borderRadius:12,cursor:"pointer"},
  hd:{display:"flex",alignItems:"center",justifyContent:"space-between",padding:"12px 14px",borderBottom:"1px solid rgba(100,116,139,0.1)",background:"rgba(12,12,30,0.85)",backdropFilter:"blur(10px)",position:"sticky",top:0,zIndex:10},
  ht:{fontSize:14,fontWeight:700,margin:0,color:"#c7d2fe"},
  bk:{background:"none",border:"none",color:"#818cf8",fontSize:12,fontWeight:600,cursor:"pointer",padding:"4px 6px"},
  tb:{padding:"9px 18px",fontSize:12,fontWeight:700,background:"rgba(99,102,241,0.1)",border:"1px solid rgba(99,102,241,0.2)",borderRadius:10,color:"#c7d2fe",cursor:"pointer"},
  inp:{flex:1,padding:"13px 14px",fontSize:18,fontWeight:700,background:"rgba(15,15,35,0.4)",border:"1px solid rgba(100,116,139,0.18)",borderRadius:12,color:"#e2e8f0",outline:"none"},
};

if(typeof document!=="undefined"&&!document.getElementById("ft-s")){
  const s=document.createElement("style");s.id="ft-s";
  s.textContent=`@keyframes pulse{0%,100%{opacity:1}50%{opacity:.5}}*{box-sizing:border-box;-webkit-tap-highlight-color:transparent}body{margin:0;background:#0c0c1d}button:hover{filter:brightness(1.1)}button:active{transform:scale(.97)}input[type=number]::-webkit-inner-spin-button,input[type=number]::-webkit-outer-spin-button{-webkit-appearance:none}input[type=number]{-moz-appearance:textfield}`;
  document.head.appendChild(s);
}
