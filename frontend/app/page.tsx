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

      {/* Tarjetas de Saldo */}
      <div className="max-w-6xl mx-auto grid grid-cols-1 md:grid-cols-3 gap-6 mb-10">
        {summary.map((item, index) => {
          // Usamos los nombres exactos de tu JSON de Go
          const currency = item.currency || 'USD';
          const amount = item.pending_movements_amount || 0; // Mostramos lo pendiente
          const reconciled = item.total_reconciled || 0;

          return (
            <div key={currency + index} className="bg-white p-6 rounded-2xl shadow-sm border border-gray-100">
              <div className="flex justify-between items-start mb-4">
                <div className="p-3 bg-blue-50 text-blue-600 rounded-xl">
                  <Wallet size={24} />
                </div>
                <span className="text-[10px] font-bold px-2 py-1 rounded-full bg-blue-100 text-blue-700 uppercase">
                  Conciliado: {new Intl.NumberFormat('es-CO').format(reconciled)}
                </span>
              </div>
              <p className="text-gray-500 text-sm font-medium">Pendiente ({currency})</p>
              <h3 className="text-2xl font-bold text-gray-900">
                {new Intl.NumberFormat('es-CO', { 
                  style: 'currency', 
                  currency: currency.length === 3 ? currency : 'USD' 
                }).format(amount)}
              </h3>
            </div>
          );
        })}
      </div>

      {/* Sección de IA (Placeholder para el siguiente paso) */}
      <div className="max-w-6xl mx-auto bg-gradient-to-r from-blue-600 to-indigo-700 rounded-2xl p-8 text-white flex flex-col md:flex-row items-center justify-between shadow-lg">
        <div className="mb-6 md:mb-0">
          <div className="flex items-center gap-2 mb-2">
            <BrainCircuit size={24} />
            <span className="font-semibold uppercase tracking-wider text-sm opacity-80">Módulo de IA Activo</span>
          </div>
          <h2 className="text-2xl font-bold">Categorizador Inteligente</h2>
          <p className="opacity-90 max-w-md mt-2">Prueba el servicio de IA para clasificar automáticamente tus movimientos bancarios.</p>
        </div>
        <div className="bg-white/10 p-4 rounded-xl backdrop-blur-md border border-white/20">
            <button className="bg-white text-blue-600 px-6 py-3 rounded-lg font-bold hover:bg-blue-50 transition flex items-center gap-2">
               Abrir Asistente <ArrowUpRight size={20} />
            </button>
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