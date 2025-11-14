using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Data.OleDb;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace Final_Ödevi
{
    public partial class Form2 : Form
    {
        public string a, b, c, d;
        public Form2()
        {
            InitializeComponent();
        }

        private void label4_Click(object sender, EventArgs e)
        {

        }

        private void button1_Click(object sender, EventArgs e)
        {
            string vtyolu = "Provider=Microsoft.Jet.OLEDB.4.0;Data Source=Veritabanı.mdb;Persist Security Info=False;";
            OleDbConnection baglanti = new OleDbConnection(vtyolu);
            string sorgu = "INSERT INTO Stok(uadi,Fiyat,Kar,Stok) values(@isim,@fiyat,@kar,@stok)";
            string sorgu2= "UPDATE Stok Set uadi=@uadi,Fiyat=@fiyat,Kar=@kar,Stok=@stok WHERE uadi=@uadi";
            
            Form1 frm1 = (Form1)Application.OpenForms["Form1"];
            if (frm1.flag)
            {
                baglanti.Open();
                OleDbCommand komut = new OleDbCommand(sorgu, baglanti);
                komut = new OleDbCommand(sorgu, baglanti);
                komut.Parameters.AddWithValue("@isim", textBox1.Text);
                komut.Parameters.AddWithValue("@fiyat", textBox2.Text);
                komut.Parameters.AddWithValue("@kar", textBox3.Text);
                komut.Parameters.AddWithValue("@stok", textBox4.Text);
                komut.ExecuteNonQuery();
                baglanti.Close();
                frm1.KisiEkle();
                this.Close();
                MessageBox.Show("Yeni Kayıt Eklendi..!");
            }
            else
            {
                baglanti.Open();
                OleDbCommand komut2 = new OleDbCommand(sorgu2, baglanti);
                komut2 = new OleDbCommand(sorgu2, baglanti);
                komut2.Parameters.AddWithValue("@uadi", textBox1.Text);
                komut2.Parameters.AddWithValue("@fiyat", textBox2.Text);
                komut2.Parameters.AddWithValue("@kar", textBox3.Text);
                komut2.Parameters.AddWithValue("@stok", textBox4.Text);
                komut2.ExecuteNonQuery();
                baglanti.Close();
                frm1.KisiEkle();
                this.Close();
                MessageBox.Show("Veri Güncellemesi Yapıldı");
            }
        }

        private void button2_Click(object sender, EventArgs e)
        {
            this.Close();
        }

        public void textBox1_TextChanged(object sender, EventArgs e)
        {
          
        }

        public void textBox2_TextChanged(object sender, EventArgs e)
        {

        }

        private void Form2_Load(object sender, EventArgs e)
        {
            textBox1.Text = a;
            textBox2.Text = b;
            textBox3.Text = c;
            textBox4.Text = d;
        }

        public void textBox3_TextChanged(object sender, EventArgs e)
        {

        }

        public void textBox4_TextChanged(object sender, EventArgs e)
        {

        }

        private void Form2_FormClosed(object sender, FormClosedEventArgs e)
        {
            Form1 f1 = (Form1)Application.OpenForms["Form1"];
            f1.KisiEkle();
        }


    }
}
