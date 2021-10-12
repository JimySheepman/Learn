using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows;
using System.Windows.Controls;
using System.Windows.Data;
using System.Windows.Documents;
using System.Windows.Input;
using System.Windows.Media;
using System.Windows.Media.Imaging;
using System.Windows.Shapes;

namespace lab_assingnment
{
    /// <summary>
    /// Interaction logic for Window1.xaml
    /// </summary>
    public partial class Window1 : Window
    {
        public Window1()
        {
            InitializeComponent();
        }

        private void result_button_Click(object sender, RoutedEventArgs e)
        {
            int choise_count = menu_listbox.SelectedItems.Count;
            switch (choise_count)
            {
                case 0:
                    MessageBox.Show("Only " + Convert.ToString(choise_count) + " activities. Not enough!");
                    break;
                case 1:
                    MessageBox.Show("Only " + Convert.ToString(choise_count) + " activities. Not enough!");
                    break;
                default:
                    MessageBox.Show(Convert.ToString(choise_count) + " activities! Healthy!");
                    break;
            }
        }

        private void web_button_Click(object sender, RoutedEventArgs e)
        {
            Window2 window2 = new Window2();
            window2.Show();
            this.Close();
        }

        private void back_button_Click(object sender, RoutedEventArgs e)
        {
            MainWindow window = new MainWindow();
            window.Show();
            this.Close();
        }
    }
}
