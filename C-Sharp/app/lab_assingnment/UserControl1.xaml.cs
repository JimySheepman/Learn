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
using System.Windows.Navigation;
using System.Windows.Shapes;

namespace lab_assingnment
{
    /// <summary>
    /// Interaction logic for UserControl1.xaml
    /// </summary>
    public partial class UserControl1 : UserControl
    {
        public UserControl1()
        {
            InitializeComponent();
        }

        private void calculate_button_Click(object sender, RoutedEventArgs e)
        {
            double hei = Convert.ToDouble(height_textbox.Text);
            double wei = Convert.ToDouble(weight_textbox.Text);
            double result = 0;
            result = wei / (Math.Pow(hei,2));
            if (result >= 18.5 && result<=25)
            {
                result = Math.Round(result, 2);
                result_textbox.Text = Convert.ToString(result);
                result_textbox.Background = Brushes.Green;

            }
            else
            {
                result = Math.Round(result, 2);
                result_textbox.Text = Convert.ToString(result);
                result_textbox.Background = Brushes.Red;

            }
        }
    }
}
