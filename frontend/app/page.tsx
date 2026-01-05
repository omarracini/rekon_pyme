"use client";

import { useEffect, useState } from 'react';
import { Wallet, BrainCircuit, RefreshCw, ArrowUpRight } from 'lucide-react';
import { DashboardSummary } from '@/types/banking';

export default function Home() {
  const [summary, setSummary] = useState<DashboardSummary[]>([]);
  const [loading, setLoading] = useState(true);
  const [concept, setConcept] = useState("");
  const [aiSuggestion, setAiSuggestion] = useState<any>(null);
  const [movements, setMovements] = useState<any[]>([]); // Estado para la tabla

  const fetchDashboard = async () => {
   try {
      setLoading(true);
      const resSummary = await fetch('http://localhost:8080/dashboard');
      const dataSummary = await resSummary.json();
      setSummary(dataSummary);

      const resMovs = await fetch('http://localhost:8080/movements/pending'); 
      const dataMovs = await resMovs.json();
      setMovements(dataMovs);
    } catch (error) {
      console.error("Error:", error);
    } finally {
      setLoading(false);
    }
  };

  const handleAISuggest = async () => {
    if (!concept) return;
    try {
      // Llamamos al endpoint de Go
      const res = await fetch(`http://localhost:8080/ai/suggest-category?concept=${encodeURIComponent(concept)}`);
      const data = await res.json();
      setAiSuggestion(data);
    } catch (error) {
      console.error("Error en IA:", error);
    }
  };

  useEffect(() => {
    fetchDashboard();
  }, []);

  return (
    <main className="min-h-screen bg-gray-50 p-8">
      {/* Header */}
      <div className="max-w-6xl mx-auto mb-10 flex justify-between items-center">
        <div>
          <h1 className="text-3xl font-bold text-gray-900">Rekon Pyme</h1>
          <p className="text-gray-500">Panel de Control y Conciliación</p>
        </div>
        <button 
          onClick={fetchDashboard}
          className="flex items-center gap-2 bg-white border px-4 py-2 rounded-lg hover:bg-gray-50 transition shadow-sm"
        >
          <RefreshCw size={18} className={loading ? "animate-spin" : ""} />
          Actualizar
        </button>
      </div>

      {/* Contenedor Principal Superior */}
      <div className="max-w-6xl mx-auto grid grid-cols-1 lg:grid-cols-3 gap-8 mb-10">
        
        {/* Sección Izquierda: Tarjetas de Saldo (Ocupa 2 columnas en desktop) */}
        <div className="lg:col-span-2 grid grid-cols-1 md:grid-cols-2 gap-4">
          {summary.map((item, index) => {
            const currency = item.currency || 'USD';
            const amount = item.pending_movements_amount || 0;
            const reconciled = item.total_reconciled || 0;

            return (
              <div key={currency + index} className="bg-white p-6 rounded-2xl shadow-sm border border-gray-100 flex flex-col justify-between">
                <div className="flex justify-between items-start mb-4">
                  <div className="p-2 bg-blue-50 text-blue-600 rounded-lg">
                    <Wallet size={20} />
                  </div>
                  <span className="text-[10px] font-bold px-2 py-0.5 rounded-full bg-blue-50 text-blue-600 uppercase tracking-wider">
                    Conciliado: {new Intl.NumberFormat('es-CO').format(reconciled)}
                  </span>
                </div>
                <div>
                  <p className="text-gray-400 text-xs font-semibold uppercase mb-1">{currency}</p>
                  <h3 className="text-2xl font-bold text-gray-900">
                    {new Intl.NumberFormat('es-CO', { 
                      style: 'currency', 
                      currency: currency.length === 3 ? currency : 'USD' 
                    }).format(amount)}
                  </h3>
                </div>
              </div>
            );
          })}
        </div>

        {/* Sección Derecha: Asistente de IA (Ocupa 1 columna) */}
        <div className="bg-blue-600 rounded-2xl p-6 text-white shadow-xl flex flex-col justify-between border border-blue-500 relative overflow-hidden">
          {/* Decoración de fondo */}
          <div className="absolute top-0 right-0 -mt-4 -mr-4 w-24 h-24 bg-white/10 rounded-full blur-2xl"></div>
          
          <div>
            <div className="flex items-center gap-2 mb-4">
              <BrainCircuit size={24} className="text-blue-100" />
              <h3 className="font-bold text-lg">Asistente IA</h3>
            </div>
            
            <p className="text-blue-100 text-xs mb-4 leading-relaxed">
              Ingresa un concepto para clasificarlo automáticamente.
            </p>

            <input 
              type="text" 
              placeholder="Ej: Pago AWS..."
              className="w-full p-3 rounded-xl text-gray-900 text-sm mb-3 border-none focus:ring-2 focus:ring-blue-300 outline-none shadow-inner"
              value={concept}
              onChange={(e) => setConcept(e.target.value)}
            />
            
            <button 
              onClick={handleAISuggest}
              className="w-full bg-white text-blue-600 py-2.5 rounded-xl font-bold hover:bg-blue-50 transition-all flex items-center justify-center gap-2 text-sm shadow-md active:scale-95"
            >
              Consultar IA
            </button>
          </div>

          {/* Resultado de la IA */}
          {aiSuggestion && (
            <div className="mt-4 p-3 bg-white/10 rounded-xl text-[12px] border border-white/20 animate-in fade-in zoom-in duration-300">
              <div className="flex justify-between mb-1">
                <span className="opacity-70 font-medium">Categoría:</span>
                <span className="font-bold">{aiSuggestion.category}</span>
              </div>
              <div className="w-full bg-white/20 h-1.5 rounded-full mt-2">
                <div 
                  className="bg-green-400 h-1.5 rounded-full transition-all duration-500" 
                  style={{ width: `${(aiSuggestion.confidence * 100)}%` }}
                ></div>
              </div>
              <p className="italic mt-2 opacity-80 leading-tight">"{aiSuggestion.reason}"</p>
            </div>
          )}
        </div>
      </div>

      {/* Tabla de Movimientos */}
      <div className="max-w-6xl mx-auto bg-white rounded-2xl shadow-sm border border-gray-100 overflow-hidden mt-10">
        <div className="p-6 border-b border-gray-100">
          <h3 className="text-lg font-bold text-gray-900">Movimientos Recientes</h3>
        </div>
        <div className="overflow-x-auto">
          <table className="w-full text-left">
            <thead className="bg-gray-50 text-gray-400 text-xs uppercase font-semibold">
              <tr>
                <th className="px-6 py-4">Concepto</th>
                <th className="px-6 py-4">Monto</th>
                <th className="px-6 py-4">Moneda</th>
                <th className="px-6 py-4">Estado</th>
              </tr>
            </thead>
            <tbody className="divide-y divide-gray-100">
              {movements.map((mov, index) => {
                // 1. Extraemos el valor del objeto Money (ajustado a la serialización de Go)
                // Probamos con 'Amount' (mayúscula) y 'amount' (minúscula)
                const montoReal = mov.Amount?.Amount || mov.Amount?.amount || mov.Amount?.Value || 0;
                
                // 2. Extraemos la moneda del objeto Money
                const monedaReal = mov.Amount?.Currency || mov.Amount?.currency || mov.Currency || 'USD';

                return (
                  <tr key={index} className="hover:bg-gray-50 transition text-sm">
                    <td className="px-6 py-4 font-medium text-gray-900">{mov.Concept || mov.concept}</td>
                    <td className={`px-6 py-4 font-bold ${montoReal < 0 ? 'text-red-600' : 'text-green-600'}`}>
                      {new Intl.NumberFormat('es-CO', { 
                        style: 'currency', 
                        currency: monedaReal.length === 3 ? monedaReal : 'USD' 
                      }).format(montoReal)}
                    </td>
                    <td className="px-6 py-4 text-gray-500">{monedaReal}</td>
                    <td className="px-6 py-4 text-center text-xs">
                      <span className="px-2 py-1 bg-yellow-100 text-yellow-700 rounded-md font-bold uppercase">
                        {mov.IsConciliated ? 'Conciliado' : 'Pendiente'}
                      </span>
                    </td>
                  </tr>
                );
              })}
            </tbody>
          </table>
        </div>
      </div>
    </main>
  );
}