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

namespace BookStore
{
    /// <summary>
    /// Interaction logic for Window2.xaml
    /// </summary>
    public partial class Window2 : Window
    {
        public Window2()
        {
            InitializeComponent();
        }
        public static List<string> vs = new List<string>();
        private void NextButton_Click(object sender, RoutedEventArgs e)
        {
            Window3 window3 = new Window3();
            window3.Show();
            this.Close();
        }

        private void NextCheckbox_Click(object sender, RoutedEventArgs e)
        {
            NextButton.IsEnabled = true;
        }

        private void AddButoon_Click(object sender, RoutedEventArgs e)
        {
            vs.Add(BookComboBox.Text);
            MessageBox.Show("Selected Book Added!");
        }

        private void CatagoryComboBox_SelectionChanged(object sender, SelectionChangedEventArgs e)
        {
            if (CatagoryComboBox.SelectedIndex==0)
            {
                BookComboBox.Items.Clear();
                BookComboBox.Items.Add("The Odyssey");
                BookComboBox.Items.Add("Gulliver's Travels");
                BookComboBox.Items.Add("Moby-Dick");

            }
            else if (CatagoryComboBox.SelectedIndex == 1)
            {
                BookComboBox.Items.Clear();
                BookComboBox.Items.Add("The Gambler");
                BookComboBox.Items.Add("I Capture The Castle ");
                BookComboBox.Items.Add("Brave New World");
            }
            else if (CatagoryComboBox.SelectedIndex == 2)
            {
                BookComboBox.Items.Clear();
                BookComboBox.Items.Add("The Lord of the Rings");
                BookComboBox.Items.Add("The Hitchhiker’s Guide to the Galaxy");
                BookComboBox.Items.Add("Animal Farm");

            }

        }
    }
}
