package project3;

import java.awt.*;
import java.awt.event.*;
import java.math.BigDecimal;
import javax.swing.*;
import java.util.Map;
import java.util.Map.*;
import java.util.Set;

/**
 * This class is the central controlller and GUI menager for the single window application. Its called from SongDatabase class.
 *
 * @author Brandon Gonzales 
 */
public class SongDatabasePanel extends JPanel implements ActionListener
{

    final static long serialVersionUID = 1240829475912732L;

    private MusicLibrary library;
    private JComboBox<Track> jcb;
    private boolean isAddMode = false;
    private boolean isEditMode = false;
    private int lastSelectedIndex = 0;

    private JTextField jtfItemCode;
    private JTextField jtfDescription;
    private JTextField jtfArtist;
    private JTextField jtfAlbum;
    private JTextField jtfPrice;

    private JLabel lblEmptyItemCode;
    private JLabel lblEmptyDescription;
    private JLabel lblEmptyArtist;
    private JLabel lblEmptyAlbum;
    private JLabel lblBadPrice;

    private JButton btnAdd;
    private JButton btnEdit;
    private JButton btnDelete;
    private JButton btnAccept;
    private JButton btnCancel;

    public SongDatabasePanel(MusicLibrary library)
    {

        //Local only variables
        JLabel label;
        JButton btnExit;

        this.library = library;

        //GRIDBAG LAYOUT
        this.setLayout(new GridBagLayout());
        GridBagConstraints c = new GridBagConstraints();
        c.fill = GridBagConstraints.HORIZONTAL;
        c.insets = new Insets(1,3,1,3);

        //ITEMCODE
        label = new JLabel("Item Code:", SwingConstants.RIGHT);
        c.gridwidth = 1;
        c.gridx = 0;
        c.gridy = 1;
        this.add(label, c);

        jtfItemCode = new JTextField("");
        c.gridwidth = 2;
        c.gridx = 1;
        c.gridy = 1;
        c.gridwidth = 1;
        c.anchor = GridBagConstraints.CENTER;
        this.add(jtfItemCode, c);

        lblEmptyItemCode = new JLabel("Empty", SwingConstants.LEFT);
        lblEmptyItemCode.setForeground(Color.red);
        lblEmptyItemCode.setVisible(false);
        c.gridwidth = 1;
        c.gridx = 2;
        c.gridy = 1;
        this.add(lblEmptyItemCode, c);

        //DESCRIPTION
        label = new JLabel("Description:", SwingConstants.RIGHT);
        c.gridwidth = 1;
        c.gridx = 0;
        c.gridy = 2;
        this.add(label, c);

        jtfDescription = new JTextField("");
        c.gridwidth = 2;
        c.gridx = 1;
        c.gridy = 2;
        c.gridwidth = 3;
        c.anchor = GridBagConstraints.CENTER;
        this.add(jtfDescription, c);

        lblEmptyDescription = new JLabel("Empty", SwingConstants.LEFT);
        lblEmptyDescription.setForeground(Color.red);
        lblEmptyDescription.setVisible(false);
        c.gridwidth = 1;
        c.gridx = 4;
        c.gridy = 2;
        this.add(lblEmptyDescription, c);

        //ARTIST
        label = new JLabel("Artist:", SwingConstants.RIGHT);
        c.gridwidth = 1;
        c.gridx = 0;
        c.gridy = 3;
        this.add(label, c);

        jtfArtist = new JTextField("");
        c.gridwidth = 2;
        c.gridx = 1;
        c.gridy = 3;
        c.gridwidth = 3;
        c.anchor = GridBagConstraints.CENTER;
        this.add(jtfArtist, c);

        lblEmptyArtist = new JLabel("Empty", SwingConstants.LEFT);
        lblEmptyArtist.setForeground(Color.red);
        lblEmptyArtist.setVisible(false);
        c.gridwidth = 1;
        c.gridx = 4;
        c.gridy = 3;
        this.add(lblEmptyArtist, c);

        //ALBUM
        label = new JLabel("Album:", SwingConstants.RIGHT);
        c.gridwidth = 1;
        c.gridx = 0;
        c.gridy = 4;
        this.add(label, c);

        jtfAlbum = new JTextField("");
        c.gridwidth = 2;
        c.gridx = 1;
        c.gridy = 4;
        c.gridwidth = 3;
        c.anchor = GridBagConstraints.CENTER;
        this.add(jtfAlbum, c);

        lblEmptyAlbum = new JLabel("Empty", SwingConstants.LEFT);
        lblEmptyAlbum.setForeground(Color.red);
        lblEmptyAlbum.setVisible(false);
        c.gridwidth = 1;
        c.gridx = 4;
        c.gridy = 4;
        this.add(lblEmptyAlbum, c);

        //PRICE
        label = new JLabel("Price:", SwingConstants.RIGHT);
        c.gridwidth = 1;
        c.gridx = 0;
        c.gridy = 5;
        this.add(label, c);

        jtfPrice = new JTextField("");
        c.gridwidth = 2;
        c.gridx = 1;
        c.gridy = 5;
        c.gridwidth = 1;
        c.anchor = GridBagConstraints.CENTER;
        this.add(jtfPrice, c);

        lblBadPrice = new JLabel("Bad Number", SwingConstants.LEFT);
        lblBadPrice.setForeground(Color.red);
        lblBadPrice.setVisible(false);
        c.gridwidth = 1;
        c.gridx = 2;
        c.gridy = 5;
        this.add(lblBadPrice, c);

        //BUTTONS
        c.ipady = 8;
        c.gridy = 6;
        c.fill = GridBagConstraints.HORIZONTAL;
        c.gridwidth = 1;

        btnAdd = new JButton("Add");
        c.weightx = 0.5;
        c.gridx = 0;
        c.insets = new Insets(50,20,1,5);
        this.add(btnAdd, c);
        c.insets = new Insets(50,5,1,5);

        btnEdit = new JButton("Edit");
        c.weightx = 0.5;
        c.gridx = 1;
        this.add(btnEdit, c);

        btnDelete = new JButton("Delete");
        c.weightx = 0.5;
        c.gridx = 2;
        this.add(btnDelete, c);

        btnAccept = new JButton("Accept");
        c.weightx = 0.5;
        c.gridx = 3;
        this.add(btnAccept, c);

        btnCancel = new JButton("Cancel");
        c.weightx = 0.5;
        c.gridx = 4;
        c.insets = new Insets(50,5,1,20);
        this.add(btnCancel, c);

        btnExit = new JButton("Exit");
        c.weightx = 0.5;
        c.gridx = 2;
        c.gridy = 7;
        c.insets = new Insets(1,5,1,5);
        this.add(btnExit, c);

        //SONG
        label = new JLabel("Select Song:", SwingConstants.RIGHT);
        c.gridx = 0;
        c.gridy = 0;
        this.add(label, c);

        jcb = new JComboBox<Track>();
        c.weightx = 1;
        c.fill = GridBagConstraints.HORIZONTAL;
        c.gridx = 1;
        c.gridy = 0;
        c.gridwidth = 3;
        this.add(jcb, c);

        setEditMode(false);
        jtfItemCode.setEnabled(false);

        //Listeners
        btnAdd.addActionListener(this);
        btnEdit.addActionListener(this);
        btnDelete.addActionListener(this);
        btnAccept.addActionListener(this);
        btnCancel.addActionListener(this);
        btnExit.addActionListener(this);
        jcb.addActionListener(this);

        //Override render to only display track name from track model object (not object id).
        jcb.setRenderer(new DefaultListCellRenderer() {
            final static long serialVersionUID = 2799829491812710L;

            @Override
            public Component getListCellRendererComponent(JList<?> list,
                                                          Object value,
                                                          int index,
                                                          boolean isSelected,
                                                          boolean cellHasFocus) {

                //This is the important part we seek to override
                Track track = (Track) value;
                if (track == null)
                {
                    value = "";
                }
                else
                {
                    value = track.getDescription();
                }

                return super.getListCellRendererComponent(list, value,
                        index, isSelected, cellHasFocus);
            }
        });

        refreshSongComboBox();

    }

    /**
     * This method sets the edit mode when the edit button is pressed. Buttons are enabled or disabled per spec.
     * @param isEditMode (boolean) if edit mode is entered
     */
    public void setEditMode(boolean isEditMode)
    {
        this.isEditMode = isEditMode;

        if (isEditMode)
        {
            jcb.setEnabled(false);
            setEnabledAllTextFields(true, false);
            btnAccept.setEnabled(true);
            btnCancel.setEnabled(true);
            btnAdd.setEnabled(false);
            btnEdit.setEnabled(false);
            btnDelete.setEnabled(false);
        }
        else
        {
            jcb.setEnabled(true);
            setEnabledAllTextFields(false, false);
            btnAccept.setEnabled(false);
            btnCancel.setEnabled(false);
            btnAdd.setEnabled(true);
            btnEdit.setEnabled(true);
            btnDelete.setEnabled(true);
        }

    }

    /**
     * This method sets the add mode when the edit button is pressed. Buttons are enabled or disabled per spec.
     * @param isAddMode (boolean) if add mode is entered
     */
    public void setAddMode(boolean isAddMode)
    {
        this.isAddMode = isAddMode;

        if (isAddMode)
        {
            jcb.setEnabled(false);
            setEnabledAllTextFields(true, true);
            jtfItemCode.setEnabled(true);
            btnEdit.setEnabled(false);
            btnDelete.setEnabled(false);
            btnAdd.setEnabled(false);
            btnAccept.setEnabled(true);
            btnCancel.setEnabled(true);
        }
        else
        {
            jcb.setEnabled(true);
            setEnabledAllTextFields(false, false);
            jtfItemCode.setEnabled(false);
            btnEdit.setEnabled(true);
            btnDelete.setEnabled(true);
            btnAdd.setEnabled(true);
            btnAccept.setEnabled(false);
            btnCancel.setEnabled(false);
        }

    }

    /**
     * This is the action handler for all controls. It listens to the actionevent and determines what kind of control was changed.
     *
     * @param e (ActionEvent) of the GUI element.
     */
    public void actionPerformed(ActionEvent e)
    {
        //Case: this is a button
        if (e.getSource() instanceof JButton)
        {

            switch (e.getActionCommand())
            {
                case "Delete":

                    //Remove from colleciton index is the hashCode
                    library.removeAtIndex(jcb.getSelectedItem().hashCode());
                    refreshSongComboBox();


                    //Special Case for last item deleted
                    if (jcb.getItemCount() <= 0)
                    {
                        btnDelete.setEnabled(false);
                        btnEdit.setEnabled(false);
                        jcb.setEnabled(false);
                        setEnabledAllTextFields(false, true);
                    }
                    break;
                case "Edit":
                    lastSelectedIndex = jcb.getSelectedIndex();     //Save last selected index
                    setAddMode(false);
                    setEditMode(true);
                    break;
                case "Add":
                    lastSelectedIndex = jcb.getSelectedIndex();     //Save last selected index
                    setEditMode(false);
                    setAddMode(true);
                    break;
                case "Cancel":
                    setEditMode(false);
                    setAddMode(false);
                    refreshSongComboBox();
                    jcb.setSelectedIndex(lastSelectedIndex);        //Goto last selected index
                    break;
                case "Accept":
                    Track saveTrack;

                    if (isEditMode)
                    {
                        saveTrack = (Track) jcb.getSelectedItem();
                    }
                    else
                    {
                        saveTrack = new Track();
                    }

                    //saveForm method does the saving in the collection
                    if (saveForm(saveTrack))
                    {

                        refreshSongComboBox();

                        //Goto last item in combobox which is just added item.
                        if (isAddMode)
                        {
                            jcb.setSelectedIndex(jcb.getItemCount() - 1);
                        }
                        else if (isEditMode)
                        {
                            jcb.setSelectedIndex(lastSelectedIndex);
                        }

                        setEditMode(false);
                        setAddMode(false);
                    }
                    break;
                case "Exit":
                    library.saveLibraryToDisk();            //Always save to disk.
                    System.exit(0);                         //Bye bye
                    break;
                default:
            }

        }
        else if (e.getSource() instanceof JComboBox)        //Combobox change
        {
            Track track = (Track) jcb.getSelectedItem();

            if (track != null)
            {
                refreshTextFields(track);
            }

        }

    }

    /**
     * This method refreshed the text fields based on a track passed in. Usually happens on combo box change.
     *
     * @param track (Track) track object to get info from.
     */
    public void refreshTextFields(Track track)
    {
        jtfItemCode.setText(track.getItemCode());
        jtfDescription.setText(track.getDescription());
        jtfArtist.setText(track.getArtist());
        jtfAlbum.setText(track.getAlbum());
        jtfPrice.setText(track.getPrice().setScale(2, BigDecimal.ROUND_HALF_UP).toString());
    }

    /**
     * This method enables and disables and clears the text fields based on two booleans.
     *
     * @param enabled (boolean) enable the editable fields.
     * @param clear (boolean) clear the fields including the item code.
     */
    public void setEnabledAllTextFields(boolean enabled, boolean clear)
    {
        jtfDescription.setEnabled(enabled);
        jtfArtist.setEnabled(enabled);
        jtfAlbum.setEnabled(enabled);
        jtfPrice.setEnabled(enabled);

        if (clear)
        {
            jcb.removeAllItems();
            jtfItemCode.setText("");
            jtfDescription.setText("");
            jtfArtist.setText("");
            jtfAlbum.setText("");
            jtfPrice.setText("");
        }

    }

    /**
     * This method saves the text field values.
     *
     * @param track (Track) the track that is to be replaced on save. A new track is passed in if its an "Add"
     * @return (boolean) success or not on the save.
     */
    public boolean saveForm(Track track)
    {

        boolean process = true;                     //If true after all checks, all is good
        BigDecimal price = new BigDecimal("-1");    //price is set here to avoid warnings.

        //Lots of one-by-one value checks

        if (jtfItemCode.getText().isEmpty())
        {
            lblEmptyItemCode.setVisible(true);
            process = false;
        }
        else
        {
            lblEmptyItemCode.setVisible(false);
        }

        if (jtfDescription.getText().isEmpty())
        {
            lblEmptyDescription.setVisible(true);
            process = false;
        }
        else
        {
            lblEmptyDescription.setVisible(false);
        }

        if (jtfArtist.getText().isEmpty())
        {
            lblEmptyArtist.setVisible(true);
            process = false;
        }
        else
        {
            lblEmptyArtist.setVisible(false);
        }

        if (jtfAlbum.getText().isEmpty())
        {
            lblEmptyAlbum.setVisible(true);
            process = false;
        }
        else
        {
            lblEmptyAlbum.setVisible(false);
        }

        try
        {
            price = new BigDecimal(jtfPrice.getText());
            lblBadPrice.setVisible(false);
        }
        catch (NumberFormatException e)
        {
            lblBadPrice.setVisible(true);
            process = false;
        }

        //All checks passed, now write.
        if (process)
        {
                track.setItemCode(jtfItemCode.getText());
                track.setDescription(jtfDescription.getText());
                track.setArtist(jtfArtist.getText());
                track.setAlbum(jtfAlbum.getText());
                track.setPrice(price);

                library.getLibrary().put(System.identityHashCode(track), track); //Write the track to the collection.
                return true;
        }

        return false;       //Something is wrong, keep trying.

    }

    /**
     * This method clears and reloads the combo box. If empty collection, it will disable some controls like combo, edit and delete buttons.
     */
    public void refreshSongComboBox()
    {
        jcb.removeAllItems();

        Set<Entry<Integer, Track>> set = library.getLibrary().entrySet();

        for(Map.Entry<Integer, Track> track : set)
        {
            jcb.addItem(track.getValue());
        }

        if (jcb.getItemCount() <= 0)
        {
            jcb.setEnabled(false);
            btnEdit.setEnabled(false);
            btnDelete.setEnabled(false);
        }

    }
}
