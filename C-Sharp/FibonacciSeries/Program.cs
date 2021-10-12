using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading;
using System.Threading.Tasks;

namespace FibonacciSeries
{
    class Program
    {
        class Container
        {
            public static List<int> fibseries = new List<int>();
        }

        public static void CreateFibSeries()
        {
            int i;
            Container.fibseries.Add(0);
            Container.fibseries.Add(1);
            for (i = 2; i < 10; ++i)
                Container.fibseries.Insert(i, Container.fibseries[i - 1] + Container.fibseries[i - 2]);
        }

        public static void thread1()
        {
            CreateFibSeries();
            Console.Write("Fibonacci Series:");
            for (int i = 0; i < Container.fibseries.Count; i++)
            {
                Console.Write(Container.fibseries[i]+" ");
                Thread.Sleep(500);
            }
            Console.WriteLine();
            Container.fibseries.Clear();
        }

        public static void thread2()
        {
            CreateFibSeries();
            Console.Write("Reversed Series:");
            for (int i = Container.fibseries.Count-1; i >= 0 ; i--)
            {
                Console.Write(Container.fibseries[i] + " ");
                Thread.Sleep(500);
            }
            Console.ReadLine();
        }

        static void Main(string[] args)
        {
            Thread t1 = new Thread(thread1);
            Thread t2 = new Thread(thread2);
            t1.Start();
            Thread.Sleep(10000);
            t2.Start();

        }

    }
}
