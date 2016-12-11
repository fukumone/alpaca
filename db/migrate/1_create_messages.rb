class CreateMessages < ActiveRecord::Migration
  def change
    create_table :messages do |t|
      t.string :name
      t.string :title
      t.string :body
      t.references :title, null: false

      t.timestamps
    end
  end
end
