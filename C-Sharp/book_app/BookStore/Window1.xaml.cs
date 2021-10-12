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
    /// Interaction logic for Window1.xaml
    /// </summary>
    public partial class Window1 : Window
    {
        public Window1()
        {
            InitializeComponent();
        }

        private void LoginButton_Click(object sender, RoutedEventArgs e)
        {
            Window2 window2 = new Window2();
            if (UsernameTxtbx.Text.ToString().Equals("admin") && PasswordBx.Password.Equals("admin") )
            {
                window2.Show();
                this.Close();
            }
            else
            {
                MessageBox.Show("Wrong password or username..!!");
            }
        }
    }
}
